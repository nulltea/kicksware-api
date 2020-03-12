package model

import (
	"errors"
	errs "github.com/pkg/errors"
	"github.com/rs/xid"
	"gopkg.in/dealancer/validate.v2"
	"time"
)

var (
	ErrProductNotFound = errors.New("sneaker product Not Found")
	ErrProductInvalid  = errors.New("sneaker product Invalid")
)

type sneakerProductService struct {
	sneakerProductRepo SneakerProductRepository
}

func NewBlogService(sneakerProductRepo SneakerProductRepository) SneakerProductService {
	return &sneakerProductService{
		sneakerProductRepo,
	}
}

func (r *sneakerProductService) Find(code string) (*SneakerProduct, error) {
	return r.sneakerProductRepo.Find(code)
}

func (r *sneakerProductService) Store(sneakerProduct *SneakerProduct) error {
	if err := validate.Validate(sneakerProduct); err != nil {
		return errs.Wrap(ErrProductInvalid, "service.sneakerProductRepo.Store")
	}
	sneakerProduct.UniqueId = xid.New().String()
	sneakerProduct.AddedAt = time.Now()
	return r.sneakerProductRepo.Store(sneakerProduct)
}
