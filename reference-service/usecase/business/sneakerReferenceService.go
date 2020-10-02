package business

import (
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/xid"
	"go.kicksware.com/api/service-common/api/rest"
	"go.kicksware.com/api/service-common/core"
	"gopkg.in/dealancer/validate.v2"

	"go.kicksware.com/api/service-common/core/meta"

	"go.kicksware.com/api/reference-service/core/model"
	"go.kicksware.com/api/reference-service/core/repo"
	"go.kicksware.com/api/reference-service/core/service"
	"go.kicksware.com/api/reference-service/env"
)

var (
	ErrReferenceNotFound = errors.New("sneaker reference Not Found")
	ErrReferenceNotValid = errors.New("sneaker reference Not Valid")
	uniqueIdFieldName = "uniqueid"
)

type referenceService struct {
	repo repo.SneakerReferenceRepository
	serviceConfig env.ServiceConfig
	communicator core.InnerCommunicator
}

func NewSneakerReferenceService(sneakerReferenceRepo repo.SneakerReferenceRepository, auth core.AuthService, config env.ServiceConfig) service.SneakerReferenceService {
	return &referenceService {
		sneakerReferenceRepo,
		config,
		rest.NewCommunicator(auth, config.Common),
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
	s.handleForeignSubquery(&query, params)
	refs, err = s.repo.FetchQuery(query, params)
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
	s.handleForeignSubquery(&query, params)
	return s.repo.Count(query, params)
}

func (s *referenceService) handleForeignSubquery(query *meta.RequestQuery, params *meta.RequestParams) (ok bool) {
	foreignKeys := make([]string, 0)
	ok = false
	_query := *query
	for key := range _query {
		if strings.Contains(key, "*/") {
			service := strings.TrimLeft(key, "*/");
			endpoint := path.Join(service, "query")
			subs := make([]*struct{
				ReferenceID string `json:"ReferenceID"`
			}, 0)
			if err := s.communicator.PostMessage(endpoint, _query[key], &subs, params); err == nil {
				for _, doc := range subs {
					foreignKeys = append(foreignKeys, doc.ReferenceID)
				}
				ok = true
			}
			delete(_query, key)
		}
	}
	appendInCondition(query, foreignKeys)
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
