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

	"github.com/timoth-y/kicksware-api/beta-service/core/meta"
	"github.com/timoth-y/kicksware-api/beta-service/core/model"
	"github.com/timoth-y/kicksware-api/beta-service/core/repo"
	"github.com/timoth-y/kicksware-api/beta-service/core/service"
	"github.com/timoth-y/kicksware-api/beta-service/env"
)

var (
	ErrBetaNotFound = errors.New("beta Not Found")
	ErrBetaNotValid = errors.New("beta Not Valid")
)

type betaService struct {
	repo repo.BetaRepository
	serviceConfig env.CommonConfig
}

func NewBetaService(betaRepo repo.BetaRepository, serviceConfig env.CommonConfig) service.BetaService {
	return &betaService{
		betaRepo,
		serviceConfig,
	}
}

func (s *betaService) FetchOne(code string, params *meta.RequestParams) (*model.Beta, error) {
	return s.repo.FetchOne(code, params)
}

func (s *betaService) Fetch(codes []string, params *meta.RequestParams) ([]*model.Beta, error) {
	return s.repo.Fetch(codes, params)
}

func (s *betaService) FetchAll(params *meta.RequestParams) ([]*model.Beta, error) {
	return s.repo.FetchAll(params)
}

func (s *betaService) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) (refs []*model.Beta, err error) {
	foreignKeys, is := s.handleForeignSubquery(query)
	refs, err = s.repo.FetchQuery(query, params)
	if err == nil && is {
		refs = funk.Filter(refs, func(ref *model.Beta) bool {
			return funk.Contains(foreignKeys, ref.UniqueID)
		}).([]*model.Beta)
	}
	return
}

func (s *betaService) StoreOne(beta *model.Beta) error {
	if err := validate.Validate(beta); err != nil {
		return errors.Wrap(ErrBetaNotValid, "service.repo.Store")
	}
	beta.UniqueID = xid.New().String()
	return s.repo.StoreOne(beta)
}

func (s *betaService) Store(betas []*model.Beta) error {
	for _, beta := range betas {
		beta.UniqueID = xid.New().String()
	}
	return s.repo.Store(betas)
}

func (s *betaService) Modify(beta *model.Beta) error {
	return s.repo.Modify(beta)
}

func (s *betaService) CountAll() (int, error) {
	return s.repo.CountAll()
}

func (s *betaService) Count(query meta.RequestQuery, params *meta.RequestParams) (int, error) {
	foreignKeys, is := s.handleForeignSubquery(query); if is {
		refs, err := s.repo.FetchQuery(query, params)
		if err == nil && is {
			refs = funk.Filter(refs, func(ref *model.Beta) bool {
				return funk.Contains(foreignKeys, ref.UniqueID)
			}).([]*model.Beta)
		}
		return len(refs), nil
	}
	return s.repo.Count(query, params)
}

func (s *betaService) handleForeignSubquery(query map[string]interface{}) (foreignKeys []string, is bool) {
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

func (s *betaService) postOnForeignService(service string, query interface{}) (keys []string, err error) {
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