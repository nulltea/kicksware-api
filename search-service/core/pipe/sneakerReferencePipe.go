package pipe

import (
	 "go.kicksware.com/api/reference-service/core/model"

	"go.kicksware.com/api/service-common/core/meta"
)

type SneakerReferencePipe interface {
	FetchOne(uniqueId string) (*model.SneakerReference, error)
	Fetch(uniqueId []string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	FetchAll(params *meta.RequestParams) ([]*model.SneakerReference, error)
	FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.SneakerReference, error)
}