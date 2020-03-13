package service

import model "product-service/core/model"


type SneakerProductService interface {
	RetrieveOne(code string) (*model.SneakerProduct, error)
	Retrieve(codes []string) ([]*model.SneakerProduct, error)
	RetrieveAll() ([]*model.SneakerProduct, error)
	RetrieveQuery(query interface{}) ([]*model.SneakerProduct, error)
	Store(sneakerProduct *model.SneakerProduct) error
	Modify(sneakerProduct *model.SneakerProduct) error
	Remove(sneakerProduct *model.SneakerProduct) error
}