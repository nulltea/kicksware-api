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

	"reference-service/core/meta"
	"reference-service/core/model"
	"reference-service/core/repo"
	"reference-service/core/service"
	"reference-service/env"
)

var (
	ErrReferenceNotFound = errors.New("sneaker reference Not Found")
	ErrReferenceNotValid = errors.New("sneaker reference Not Valid")
)

type referenceService struct {
	repo repo.SneakerReferenceRepository
	serviceConfig env.CommonConfig
}

func NewSneakerReferenceService(sneakerReferenceRepo repo.SneakerReferenceRepository, serviceConfig env.CommonConfig) service.SneakerReferenceService {
	return &referenceService {
		sneakerReferenceRepo,
		serviceConfig,
	}
}

func (s *referenceService) FetchOne(code string, params *meta.RequestParams) (*model.SneakerReference, error) {
	return s.repo.FetchOne(code, params)
}

func (s *referenceService) Fetch(codes []string, params *meta.RequestParams) ([]*model.SneakerReference, error) {
	return s.repo.Fetch(codes, params)
}

func (s *referenceService) FetchAll(params *meta.RequestParams) ([]*model.SneakerReference, error) {
	return s.repo.FetchAll(params)
}

func (s *referenceService) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) (refs []*model.SneakerReference, err error) {
	foreignKeys, is := s.handleForeignSubquery(query)
	refs, err = s.repo.FetchQuery(query, params)
	if err == nil && is {
		refs = funk.Filter(refs, func(ref *model.SneakerReference) bool {
			return funk.Contains(foreignKeys, ref.UniqueId)
		}).([]*model.SneakerReference)
	}
	return
}

func (s *referenceService) StoreOne(sneakerReference *model.SneakerReference) error {
	if err := validate.Validate(sneakerReference); err != nil {
		return errors.Wrap(ErrReferenceNotValid, "service.repo.Store")
	}
	sneakerReference.UniqueId = xid.New().String()
	return s.repo.StoreOne(sneakerReference)
}

func (s *referenceService) Store(sneakerReferences []*model.SneakerReference) error {
	for _, sneakerReference := range sneakerReferences {
		sneakerReference.UniqueId = xid.New().String()
	}
	return s.repo.Store(sneakerReferences)
}

func (s *referenceService) Modify(sneakerReference *model.SneakerReference) error {
	return s.repo.Modify(sneakerReference)
}

func (s *referenceService) CountAll() (int, error) {
	return s.repo.CountAll()
}

func (s *referenceService) Count(query meta.RequestQuery, params *meta.RequestParams) (int, error) {
	foreignKeys, is := s.handleForeignSubquery(query); if is {
		refs, err := s.repo.FetchQuery(query, params)
		if err == nil && is {
			refs = funk.Filter(refs, func(ref *model.SneakerReference) bool {
				return funk.Contains(foreignKeys, ref.UniqueId)
			}).([]*model.SneakerReference)
		}
		return len(refs), nil
	}
	return s.repo.Count(query, params)
}

func (s *referenceService) handleForeignSubquery(query map[string]interface{}) (foreignKeys []string, is bool) {
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

func (s *referenceService) postOnForeignService(service string, query interface{}) (keys []string, err error) {
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