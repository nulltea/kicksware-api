package service

import (
	"go.kicksware.com/api/shared/core/meta"

	"go.kicksware.com/api/services/references/core/model"
)

type SneakerReferenceService interface {
	FetchOne(code string, params *meta.RequestParams) (*model.SneakerReference, error)
	Fetch(codes []string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	FetchAll(params *meta.RequestParams) ([]*model.SneakerReference, error)
	FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.SneakerReference, error)
	StoreOne(sneakerReference *model.SneakerReference) error
	Store(sneakerReference []*model.SneakerReference) error
	Modify(sneakerReferences *model.SneakerReference) error
	Count(query meta.RequestQuery, params *meta.RequestParams) (int, error)
	CountAll() (int, error)
}
