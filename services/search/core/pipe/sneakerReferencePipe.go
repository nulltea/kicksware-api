package pipe

import (
	 "go.kicksware.com/api/services/references/core/model"

	"go.kicksware.com/api/shared/core/meta"
)

type SneakerReferencePipe interface {
	FetchOne(uniqueId string) (*model.SneakerReference, error)
	Fetch(uniqueId []string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	FetchAll(params *meta.RequestParams) ([]*model.SneakerReference, error)
	FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.SneakerReference, error)
}
