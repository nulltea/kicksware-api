package business

import (
	"errors"
	errs "github.com/pkg/errors"
	"github.com/rs/xid"
	"gopkg.in/dealancer/validate.v2"
	"product-service/core/model"
	"product-service/core/repo"
	"product-service/core/service"
	"time"
)

var (
	ErrProductNotFound = errors.New("sneaker product Not Found")
	ErrProductInvalid  = errors.New("sneaker product Invalid")
)

type SneakerProductService struct {
	sneakerProductRepo repo.SneakerProductRepository
}

func NewSneakerProductService(sneakerProductRepo repo.SneakerProductRepository) service.SneakerProductService {
	return &SneakerProductService{
		sneakerProductRepo,
	}
}

func (r *SneakerProductService) RetrieveOne(code string) (*model.SneakerProduct, error) {
	return r.sneakerProductRepo.RetrieveOne(code)
}

func (r *SneakerProductService) Retrieve(codes []string) ([]*model.SneakerProduct, error) {
	return r.sneakerProductRepo.Retrieve(codes)
}

func (r *SneakerProductService) RetrieveAll() ([]*model.SneakerProduct, error) {
	return r.sneakerProductRepo.RetrieveAll()
}

func (r *SneakerProductService) RetrieveQuery(query interface{}) ([]*model.SneakerProduct, error) {
	return r.sneakerProductRepo.RetrieveQuery(query)
}

func (r *SneakerProductService) Store(sneakerProduct *model.SneakerProduct) error {
	if err := validate.Validate(sneakerProduct); err != nil {
		return errs.Wrap(ErrProductInvalid, "service.sneakerProductRepo.Store")
	}
	sneakerProduct.UniqueId = xid.New().String()
	sneakerProduct.AddedAt = time.Now()
	return r.sneakerProductRepo.Store(sneakerProduct)
}

func (r *SneakerProductService) Modify(sneakerProduct *model.SneakerProduct) error {
	return r.sneakerProductRepo.Modify(sneakerProduct)
}

func (r *SneakerProductService) Remove(sneakerProduct *model.SneakerProduct) error {
	return r.sneakerProductRepo.Remove(sneakerProduct)
}

