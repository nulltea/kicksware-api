package repo

import (
	"github.com/timoth-y/kicksware-platform/middleware-service/reference-service/core/meta"
	"github.com/timoth-y/kicksware-platform/middleware-service/reference-service/core/model"
)


type SneakerReferenceRepository interface {
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
