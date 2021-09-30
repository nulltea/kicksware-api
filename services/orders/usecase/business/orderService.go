package business

import (
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/xid"
	"go.kicksware.com/api/services/search/core/pipe"
	"go.kicksware.com/api/shared/api/rest"
	"go.kicksware.com/api/shared/core"
	"gopkg.in/dealancer/validate.v2"

	"go.kicksware.com/api/shared/core/meta"

	"go.kicksware.com/api/services/orders/core/model"
	"go.kicksware.com/api/services/orders/core/repo"
	"go.kicksware.com/api/services/orders/core/service"
	"go.kicksware.com/api/services/orders/env"
)

var (
	ErrOrderNotFound = errors.New("order Not Found")
	ErrOrderNotValid = errors.New("order Not Valid")
	uniqueIdFieldName = "unique_id"
)

type orderService struct {
	repo          repo.OrderRepository
	pipe          pipe.SneakerReferencePipe
	serviceConfig env.ServiceConfig
	communicator  core.InnerCommunicator
}

func NewOrderService(orderRepo repo.OrderRepository, pipe pipe.SneakerReferencePipe, auth core.AuthService, config env.ServiceConfig) service.OrderService {
	return &orderService{
		orderRepo,
		pipe,
		config,
		rest.NewCommunicator(auth, config.Common),
	}
}

func (s *orderService) FetchOne(code string, params *meta.RequestParams) (*model.Order, error) {
	return s.repo.FetchOne(code, params)
}

func (s *orderService) Fetch(codes []string, params *meta.RequestParams) ([]*model.Order, error) {
	return s.repo.Fetch(codes, params)
}

func (s *orderService) FetchAll(params *meta.RequestParams) ([]*model.Order, error) {
	return s.repo.FetchAll(params)
}

func (s *orderService) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) (refs []*model.Order, err error) {
	s.handleForeignSubquery(&query, params)
	refs, err = s.repo.FetchQuery(query, params)
	return
}

func (s *orderService) StoreOne(order *model.Order) error {
	if err := validate.Validate(order); err != nil {
		return errors.Wrap(ErrOrderNotValid, "service.repo.Store")
	}
	order.UniqueID = xid.New().String()
	if ref, err := s.pipe.FetchOne(order.ReferenceID); err == nil {
		order.Price = ref.Price
		if url := ref.StadiumUrl; len(url) != 0 {
			order.SourceURL = url
		} else if url := ref.GoatUrl; len(url) != 0  {
			order.SourceURL = url
		} else {
			order.SourceURL = fmt.Sprintf("http://kicksware.com/shop/references/%v", ref.UniqueId)
		}
		order.Price = ref.Price
	}
	order.Status = model.Draft
	order.AddedAt = time.Now()
	return s.repo.StoreOne(order)
}

func (s *orderService) Modify(order *model.Order) error {
	return s.repo.Modify(order)
}

func (s *orderService) Remove(code string) error {
	return s.repo.Remove(code)
}

func (s *orderService) Count(query meta.RequestQuery, params *meta.RequestParams) (int, error) {
	s.handleForeignSubquery(&query, params)
	return s.repo.Count(query, params)
}

func (s *orderService) CountAll() (int, error) {
	return s.repo.CountAll()
}

func (s *orderService) handleForeignSubquery(query *meta.RequestQuery, params *meta.RequestParams) (ok bool) {
	foreignKeys := make([]string, 0)
	ok = false
	_query := *query
	for key := range _query {
		if strings.Contains(key, "*/") {
			service := strings.TrimLeft(key, "*/");
			endpoint := path.Join(service, "query")
			subs := make([]*struct{
				OrderID string `json:"OrderID"`
			}, 0)
			if err := s.communicator.PostMessage(endpoint, _query[key], &subs, params); err == nil {
				for _, doc := range subs {
					foreignKeys = append(foreignKeys, doc.OrderID)
				}
				ok = true
			}
			delete(_query, key)
		}
	}
	if ok {
		appendInCondition(query, foreignKeys)
	}
	return
}

func appendInCondition(query *meta.RequestQuery, keys []string) {
	in := meta.RequestQuery{ "$in": keys }
	inQuery := meta.RequestQuery{
		uniqueIdFieldName: in,
	}
	_query := *query
	if and, ok := _query["$and"]; ok {
		and = append(and.([]interface{}), inQuery)
		_query["$and"] = and
	} else if len(_query) > 0 {
		*query = meta.RequestQuery{
			"$and": []interface{}{ inQuery, _query },
		}
	} else {
		_query[uniqueIdFieldName] = in
	}
}
