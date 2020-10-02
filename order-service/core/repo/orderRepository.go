package repo

import (
	"go.kicksware.com/api/service-common/core/meta"
	"go.kicksware.com/api/order-service/core/model"
)


type OrderRepository interface {
	FetchOne(code string, params *meta.RequestParams) (*model.Order, error)
	Fetch(codes []string, params *meta.RequestParams) ([]*model.Order, error)
	FetchAll(params *meta.RequestParams) ([]*model.Order, error)
	FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.Order, error)
	StoreOne(order *model.Order) error
	Modify(orders *model.Order) error
	Remove(code string) error
	Count(query meta.RequestQuery, params *meta.RequestParams) (int, error)
	CountAll() (int, error)
}
