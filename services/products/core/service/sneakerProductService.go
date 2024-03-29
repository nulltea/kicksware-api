package service

import (
	"go.kicksware.com/api/services/products/core/model"
	"go.kicksware.com/api/shared/core/meta"
)

type SneakerProductService interface {
	FetchOne(code string) (*model.SneakerProduct, error)
	Fetch(codes []string, params *meta.RequestParams) ([]*model.SneakerProduct, error)
	FetchAll(params *meta.RequestParams) ([]*model.SneakerProduct, error)
	FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.SneakerProduct, error)
	Store(sneakerProduct *model.SneakerProduct, params *meta.RequestParams) error
	Modify(sneakerProduct *model.SneakerProduct) error
	Replace(sneakerProduct *model.SneakerProduct) error
	Remove(code string) error
	Count(query meta.RequestQuery, params *meta.RequestParams) (int, error)
	CountAll() (int, error)
}
