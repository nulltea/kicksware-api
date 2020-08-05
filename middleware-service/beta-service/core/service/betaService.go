package service

import (
	"beta-service/core/meta"
	"beta-service/core/model"
)

type BetaService interface {
	FetchOne(code string, params *meta.RequestParams) (*model.Beta, error)
	Fetch(codes []string, params *meta.RequestParams) ([]*model.Beta, error)
	FetchAll(params *meta.RequestParams) ([]*model.Beta, error)
	FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.Beta, error)
	StoreOne(beta *model.Beta) error
	Store(beta []*model.Beta) error
	Modify(betas *model.Beta) error
	Count(query meta.RequestQuery, params *meta.RequestParams) (int, error)
	CountAll() (int, error)
}