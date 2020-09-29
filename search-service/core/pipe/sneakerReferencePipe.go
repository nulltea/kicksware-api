package pipe

import (
	 "github.com/timoth-y/kicksware-api/reference-service/core/model"

	"github.com/timoth-y/kicksware-api/service-common/core/meta"
)

type SneakerReferencePipe interface {
	FetchOne(uniqueId string) (*model.SneakerReference, error)
	Fetch(uniqueId []string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	FetchAll(params *meta.RequestParams) ([]*model.SneakerReference, error)
	FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.SneakerReference, error)
}