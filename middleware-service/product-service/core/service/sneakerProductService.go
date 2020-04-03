package service

import "product-service/core/model"

type SneakerProductService interface {
	FetchOne(code string) (*model.SneakerProduct, error)
	Fetch(codes []string) ([]*model.SneakerProduct, error)
	FetchAll() ([]*model.SneakerProduct, error)
	FetchQuery(query interface{}) ([]*model.SneakerProduct, error)
	Store(sneakerProduct *model.SneakerProduct) error
	Modify(sneakerProduct *model.SneakerProduct) error
	Replace(sneakerProduct *model.SneakerProduct) error
	Remove(code string) error
	RemoveObj(sneakerProduct *model.SneakerProduct) error
}