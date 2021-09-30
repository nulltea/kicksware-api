package repo

import (
	"go.kicksware.com/api/shared/core/meta"

	"go.kicksware.com/api/services/products/core/model"
)


type SneakerProductRepository interface {
	FetchOne(uniqueId string) (*model.SneakerProduct, error)
	Fetch(uniqueId []string, params *meta.RequestParams) ([]*model.SneakerProduct, error)
	FetchAll(params *meta.RequestParams) ([]*model.SneakerProduct, error)
	FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.SneakerProduct, error)
	Store(sneakerProduct *model.SneakerProduct) error
	Modify(sneakerProduct *model.SneakerProduct) error
	Replace(sneakerProduct *model.SneakerProduct) error
	Remove(code string) error
	Count(query meta.RequestQuery, params *meta.RequestParams) (int, error)
	CountAll() (int, error)
}
