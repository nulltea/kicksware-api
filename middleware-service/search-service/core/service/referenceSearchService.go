package service

import (
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/meta"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/model"
)

type ReferenceSearchService interface {
	Search(query string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	SearchBy(field, query string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	SearchSKU(sku string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	SearchBrand(brand string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	SearchModel(model string, params *meta.RequestParams) ([]*model.SneakerReference, error)
}