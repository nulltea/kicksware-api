package business

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/thoas/go-funk"
	"gopkg.in/dealancer/validate.v2"

	"reference-service/core/model"
	"reference-service/core/repo"
	"reference-service/core/service"
)

var (
	ErrReferenceNotFound = errors.New("sneaker reference Not Found")
	ErrReferenceNotValid = errors.New("sneaker reference Not Valid")
)

type referenceService struct {
	sneakerReferenceRepo repo.SneakerReferenceRepository
}

func NewSneakerReferenceService(sneakerReferenceRepo repo.SneakerReferenceRepository) service.SneakerReferenceService {
	return &referenceService {
		sneakerReferenceRepo,
	}
}

func (s *referenceService) FetchOne(code string) (*model.SneakerReference, error) {
	return s.sneakerReferenceRepo.FetchOne(code)
}

func (s *referenceService) Fetch(codes []string) ([]*model.SneakerReference, error) {
	return s.sneakerReferenceRepo.Fetch(codes)
}

func (s *referenceService) FetchAll() ([]*model.SneakerReference, error) {
	return s.sneakerReferenceRepo.FetchAll()
}

func (s *referenceService) FetchQuery(query map[string]interface{}) (refs []*model.SneakerReference, err error) {
	foreignKeys, sub := s.handleSubquery(query)
	refs, err = s.sneakerReferenceRepo.FetchQuery(query)
	if err == nil && sub {
		refs = funk.Filter(refs, func(ref *model.SneakerReference) bool {
			return funk.Contains(foreignKeys, ref.UniqueId)
		}).([]*model.SneakerReference)
	}
	return
}

func (s *referenceService) StoreOne(sneakerReference *model.SneakerReference) error {
	if err := validate.Validate(sneakerReference); err != nil {
		return errors.Wrap(ErrReferenceNotValid, "service.sneakerReferenceRepo.Store")
	}
	sneakerReference.UniqueId = xid.New().String()
	return s.sneakerReferenceRepo.StoreOne(sneakerReference)
}

func (s *referenceService) Store(sneakerReferences []*model.SneakerReference) error {
	for _, sneakerReference := range sneakerReferences {
		sneakerReference.UniqueId = xid.New().String()
	}
	return s.sneakerReferenceRepo.Store(sneakerReferences)
}

func (s *referenceService) Modify(sneakerReference *model.SneakerReference) error {
	return s.sneakerReferenceRepo.Modify(sneakerReference)
}

func (s *referenceService) handleSubquery(query map[string]interface{}) (foreignKeys []string, sub bool) {
	foreignKeys = make([]string, 0)
	for key := range query {
		if strings.Contains(key, "*/") {
			sub = true
			res := strings.TrimLeft(key, "*/");
			host := fmt.Sprintf("%s-service", strings.Split(res, "/")[0]);
			service := fmt.Sprintf(os.Getenv("INNER_SERVICE_FORMAT"), host, res)
			if keys, err := s.postForeignService(service, query[key]); err == nil {
				foreignKeys = append(foreignKeys, keys...)
			}
			delete(query, key)
		}
	}
	return
}

func (s *referenceService) postForeignService(service string, body interface{}) (keys []string, err error) {
	subquery, _ := json.Marshal(body)
	resp, err := http.Post(service, os.Getenv("CONTENT_TYPE"), bytes.NewBuffer(subquery))
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