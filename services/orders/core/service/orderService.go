package service

import (
	"go.kicksware.com/api/shared/core/meta"

	"go.kicksware.com/api/services/orders/core/model"
)

type OrderService interface {
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
