package business

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/thoas/go-funk"
	"gopkg.in/dealancer/validate.v2"

	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/core/meta"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/core/model"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/core/repo"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/core/service"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/env"
)

var (
	ErrOrderNotFound = errors.New("order Not Found")
	ErrOrderNotValid = errors.New("order Not Valid")
)

type orderService struct {
	repo repo.OrderRepository
	serviceConfig env.CommonConfig
}

func NewOrderService(orderRepo repo.OrderRepository, serviceConfig env.CommonConfig) service.OrderService {
	return &orderService{
		orderRepo,
		serviceConfig,
	}
}

func (s *orderService) FetchOne(code string, params meta.RequestParams) (*model.Order, error) {
	return s.repo.FetchOne(code, params)
}

func (s *orderService) Fetch(codes []string, params meta.RequestParams) ([]*model.Order, error) {
	return s.repo.Fetch(codes, params)
}

func (s *orderService) FetchAll(params meta.RequestParams) ([]*model.Order, error) {
	return s.repo.FetchAll(params)
}

func (s *orderService) FetchQuery(query meta.RequestQuery, params meta.RequestParams) (refs []*model.Order, err error) {
	foreignKeys, is := s.handleForeignSubquery(query)
	refs, err = s.repo.FetchQuery(query, params)
	if err == nil && is {
		refs = funk.Filter(refs, func(ref *model.Order) bool {
			return funk.Contains(foreignKeys, ref.UniqueID)
		}).([]*model.Order)
	}
	return
}

func (s *orderService) StoreOne(order *model.Order) error {
	if err := validate.Validate(order); err != nil {
		return errors.Wrap(ErrOrderNotValid, "service.repo.Store")
	}
	order.UniqueID = xid.New().String()
	return s.repo.StoreOne(order)
}

func (s *orderService) Store(orders []*model.Order) error {
	for _, order := range orders {
		order.UniqueID = xid.New().String()
	}
	return s.repo.Store(orders)
}

func (s *orderService) Modify(order *model.Order) error {
	return s.repo.Modify(order)
}

func (s *orderService) CountAll() (int, error) {
	return s.repo.CountAll()
}

func (s *orderService) Count(query meta.RequestQuery, params meta.RequestParams) (int, error) {
	foreignKeys, is := s.handleForeignSubquery(query); if is {
		refs, err := s.repo.FetchQuery(query, params)
		if err == nil && is {
			refs = funk.Filter(refs, func(ref *model.Order) bool {
				return funk.Contains(foreignKeys, ref.UniqueID)
			}).([]*model.Order)
		}
		return len(refs), nil
	}
	return s.repo.Count(query, params)
}

func (s *orderService) handleForeignSubquery(query map[string]interface{}) (foreignKeys []string, is bool) {
	foreignKeys = make([]string, 0)
	for key := range query {
		if strings.Contains(key, "*/") {
			is = true
			res := strings.TrimLeft(key, "*/");
			host := fmt.Sprintf("%s-service", strings.Split(res, "/")[0]);
			service := fmt.Sprintf(s.serviceConfig.InnerServiceFormat, host, res)
			if keys, err := s.postOnForeignService(service, query[key]); err == nil {
				foreignKeys = append(foreignKeys, keys...)
			}
			delete(query, key)
		}
	}
	return
}

func (s *orderService) postOnForeignService(service string, query interface{}) (keys []string, err error) {
	body, err := json.Marshal(query); if err != nil {
		return
	}
	resp, err := http.Post(service, s.serviceConfig.ContentType, bytes.NewBuffer(body))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body); if err != nil {
		return
	}

	subs := make([]map[string]interface{}, 0)
	err = json.Unmarshal(bytes, &subs); if err != nil {
		return
	}

	keys = make([]string, 0)
	for _, doc := range subs {
		if key, ok := doc["ReferenceId"]; ok {
			keys = append(keys, key.(string))
		}
	}
	return
}