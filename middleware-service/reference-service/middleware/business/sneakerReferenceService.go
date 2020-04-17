package business

import (
	"github.com/pkg/errors"
	"github.com/rs/xid"
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

func (r *referenceService) FetchOne(code string) (*model.SneakerReference, error) {
	return r.sneakerReferenceRepo.FetchOne(code)
}

func (r *referenceService) Fetch(codes []string) ([]*model.SneakerReference, error) {
	return r.sneakerReferenceRepo.Fetch(codes)
}

func (r *referenceService) FetchAll() ([]*model.SneakerReference, error) {
	return r.sneakerReferenceRepo.FetchAll()
}

func (r *referenceService) FetchQuery(query interface{}) ([]*model.SneakerReference, error) {
	return r.sneakerReferenceRepo.FetchQuery(query)
}

func (r *referenceService) StoreOne(sneakerReference *model.SneakerReference) error {
	if err := validate.Validate(sneakerReference); err != nil {
		return errors.Wrap(ErrReferenceNotValid, "service.sneakerReferenceRepo.Store")
	}
	sneakerReference.UniqueId = xid.New().String()
	return r.sneakerReferenceRepo.StoreOne(sneakerReference)
}

func (r *referenceService) Store(sneakerReferences []*model.SneakerReference) error {
	for _, sneakerReference := range sneakerReferences {
		sneakerReference.UniqueId = xid.New().String()
	}
	return r.sneakerReferenceRepo.Store(sneakerReferences)
}

func (r *referenceService) Modify(sneakerReference *model.SneakerReference) error {
	return r.sneakerReferenceRepo.Modify(sneakerReference)
}