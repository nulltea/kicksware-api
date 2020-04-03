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

type productService struct {
	sneakerProductRepo repo.SneakerProductRepository
}

func NewSneakerProductService(sneakerProductRepo repo.SneakerProductRepository) service.SneakerProductService {
	return &productService{
		sneakerProductRepo,
	}
}

func (r *productService) FetchOne(code string) (*model.SneakerProduct, error) {
	return r.sneakerProductRepo.FetchOne(code)
}

func (r *productService) Fetch(codes []string) ([]*model.SneakerProduct, error) {
	return r.sneakerProductRepo.Fetch(codes)
}

func (r *productService) FetchAll() ([]*model.SneakerProduct, error) {
	return r.sneakerProductRepo.FetchAll()
}

func (r *productService) FetchQuery(query interface{}) ([]*model.SneakerProduct, error) {
	return r.sneakerProductRepo.FetchQuery(query)
}

func (r *productService) Store(sneakerProduct *model.SneakerProduct) error {
	if err := validate.Validate(sneakerProduct); err != nil {
		return errs.Wrap(ErrProductInvalid, "service.sneakerProductRepo.Store")
	}
	sneakerProduct.UniqueId = xid.New().String()
	sneakerProduct.AddedAt = time.Now()
	return r.sneakerProductRepo.Store(sneakerProduct)
}

func (r *productService) Modify(sneakerProduct *model.SneakerProduct) error {
	return r.sneakerProductRepo.Modify(sneakerProduct)
}

func (r *productService) Replace(sneakerProduct *model.SneakerProduct) error {
	return r.sneakerProductRepo.Replace(sneakerProduct)
}

func (r *productService) Remove(code string) error {
	return r.sneakerProductRepo.Remove(code)
}

func (r *productService) RemoveObj(sneakerProduct *model.SneakerProduct) error {
	return r.sneakerProductRepo.RemoveObj(sneakerProduct)
}

