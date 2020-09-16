package service

import (
	"github.com/timoth-y/kicksware-api/reference-service/core/model"

	"github.com/timoth-y/kicksware-api/search-service/core/meta"
)

type ReferenceSearchService interface {
	Search(query string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	SearchBy(field, query string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	SearchSKU(sku string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	SearchBrand(brand string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	SearchModel(model string, params *meta.RequestParams) ([]*model.SneakerReference, error)
}