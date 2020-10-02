package business

import (
	"errors"
	"strings"
	"time"

	errs "github.com/pkg/errors"
	"github.com/rs/xid"
	"go.kicksware.com/api/service-common/api/rest"
	"go.kicksware.com/api/service-common/config"
	"go.kicksware.com/api/service-common/core"
	"gopkg.in/dealancer/validate.v2"

	"go.kicksware.com/api/service-common/core/meta"

	"go.kicksware.com/api/product-service/core/model"
	"go.kicksware.com/api/product-service/core/repo"
	"go.kicksware.com/api/product-service/core/service"
)

var (
	ErrProductNotFound = errors.New("sneaker product Not Found")
	ErrProductInvalid  = errors.New("sneaker product Invalid")
	uniqueIdFieldName = "uniqueid"
)

type productService struct {
	repo repo.SneakerProductRepository
	serviceConfig config.CommonConfig
	communicator core.InnerCommunicator
}

func NewSneakerProductService(sneakerProductRepo repo.SneakerProductRepository, auth core.AuthService, config config.CommonConfig) service.SneakerProductService {
	return &productService{
		sneakerProductRepo,
		config,
		rest.NewCommunicator(auth, config),

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
	s.handleForeignSubquery(&query, params)
	products, err = s.repo.FetchQuery(query, params)
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
	s.handleForeignSubquery(&query, params)
	return s.repo.Count(query, params)
}

func (s *productService) handleForeignSubquery(query *meta.RequestQuery, params *meta.RequestParams) (ok bool) {
	foreignKeys := make([]string, 0)
	ok = false
	_query := *query
	for key := range _query {
		if strings.Contains(key, "*/") {
			endpoint := strings.TrimLeft(key, "*/");
			subs := make([]*struct{
				ProductId string `json:"ProductId"`
			}, 0)
			if err := s.communicator.PostMessage(endpoint, _query[key], &subs, params); err == nil {
				for _, doc := range subs {
					foreignKeys = append(foreignKeys, doc.ProductId)
				}
				ok = true
			}
			delete(_query, key)
		}
	}
	if len(foreignKeys) > 0 {
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
