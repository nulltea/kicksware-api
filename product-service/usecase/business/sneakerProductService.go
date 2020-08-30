package business

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	errs "github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/thoas/go-funk"
	"gopkg.in/dealancer/validate.v2"

	"github.com/timoth-y/kicksware-platform/middleware-service/product-service/core/meta"
	"github.com/timoth-y/kicksware-platform/middleware-service/product-service/core/model"
	"github.com/timoth-y/kicksware-platform/middleware-service/product-service/core/repo"
	"github.com/timoth-y/kicksware-platform/middleware-service/product-service/core/service"
	"github.com/timoth-y/kicksware-platform/middleware-service/product-service/env"
)

var (
	ErrProductNotFound = errors.New("sneaker product Not Found")
	ErrProductInvalid  = errors.New("sneaker product Invalid")
)

type productService struct {
	repo repo.SneakerProductRepository
	serviceConfig env.CommonConfig
}

func NewSneakerProductService(sneakerProductRepo repo.SneakerProductRepository, serviceConfig env.CommonConfig) service.SneakerProductService {
	return &productService{
		sneakerProductRepo,
		serviceConfig,
	}
}

func (s *productService) FetchOne(code string) (*model.SneakerProduct, error) {
	return s.repo.FetchOne(code)
}

func (s *productService) Fetch(codes []string, params *meta.RequestParams) ([]*model.SneakerProduct, error) {
	return s.repo.Fetch(codes, params)
}

func (s *productService) FetchAll(params *meta.RequestParams) ([]*model.SneakerProduct, error) {
	return s.repo.FetchAll(params)
}

func (s *productService) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) (products[]*model.SneakerProduct, err error) {
	foreignKeys, is := s.handleForeignSubquery(query)
	products, err = s.repo.FetchQuery(query, params)
	if err == nil && is {
		products = funk.Filter(products, func(ref *model.SneakerProduct) bool {
			return funk.Contains(foreignKeys, ref.UniqueId)
		}).([]*model.SneakerProduct)
	}
	return
}

func (s *productService) Store(sneakerProduct *model.SneakerProduct, params *meta.RequestParams) error {
	if err := validate.Validate(sneakerProduct); err != nil {
		return errs.Wrap(ErrProductInvalid, "service.repo.Store")
	}
	sneakerProduct.UniqueId = xid.New().String()
	sneakerProduct.AddedAt = time.Now()
	if params != nil {
		sneakerProduct.Owner = params.UserID()
	}
	return s.repo.Store(sneakerProduct)
}

func (s *productService) Modify(sneakerProduct *model.SneakerProduct) error {
	return s.repo.Modify(sneakerProduct)
}

func (s *productService) Replace(sneakerProduct *model.SneakerProduct) error {
	return s.repo.Replace(sneakerProduct)
}

func (s *productService) Remove(code string) error {
	return s.repo.Remove(code)
}

func (s *productService) CountAll() (int, error) {
	return s.repo.CountAll()
}

func (s *productService) Count(query meta.RequestQuery, params *meta.RequestParams) (int, error) {
	foreignKeys, is := s.handleForeignSubquery(query); if is {
		products, err := s.repo.FetchQuery(query, params)
		if err == nil && is {
			products = funk.Filter(products, func(ref *model.SneakerProduct) bool {
				return funk.Contains(foreignKeys, ref.UniqueId)
			}).([]*model.SneakerProduct)
		}
		return len(products), nil
	}
	return s.repo.Count(query, params)
}

func (s *productService) handleForeignSubquery(query map[string]interface{}) (foreignKeys []string, is bool) {
	foreignKeys = make([]string, 0)
	for key := range query {
		if strings.Contains(key, "*/") {
			is = true
			res := strings.TrimLeft(key, "*/");
			host := fmt.Sprintf("%s-service", strings.Split(res, "/")[0]);
			service := fmt.Sprintf(s.serviceConfig.InnerServiceFormat, host, res)
			if keys, err := s.postForeignService(service, query[key]); err == nil {
				foreignKeys = append(foreignKeys, keys...)
			}
			delete(query, key)
		}
	}
	return
}

func (s *productService) postForeignService(service string, body interface{}) (keys []string, err error) {
	query, err := json.Marshal(body); if err != nil {
		return
	}
	resp, err := http.Post(service, s.serviceConfig.ContentType, bytes.NewBuffer(query))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	subs := make([]map[string]interface{}, 0)
	err = json.Unmarshal(bytes, &subs)
	if err != nil {
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