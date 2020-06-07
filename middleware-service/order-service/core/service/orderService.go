package service

import (
	"order-service/core/meta"
	"order-service/core/model"
)

type OrderService interface {
	FetchOne(code string, params meta.RequestParams) (*model.Order, error)
	Fetch(codes []string, params meta.RequestParams) ([]*model.Order, error)
	FetchAll(params meta.RequestParams) ([]*model.Order, error)
	FetchQuery(query meta.RequestQuery, params meta.RequestParams) ([]*model.Order, error)
	StoreOne(order *model.Order) error
	Store(order []*model.Order) error
	Modify(orders *model.Order) error
	Count(query meta.RequestQuery, params meta.RequestParams) (int, error)
	CountAll() (int, error)
}