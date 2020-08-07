package pipe

import (
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/meta"
	model "github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/model"
)

type SneakerReferencePipe interface {
	FetchOne(uniqueId string) (*model.SneakerReference, error)
	Fetch(uniqueId []string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	FetchAll(params *meta.RequestParams) ([]*model.SneakerReference, error)
	FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.SneakerReference, error)
}