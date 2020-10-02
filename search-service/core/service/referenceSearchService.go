package service

import (
	"go.kicksware.com/api/reference-service/core/model"

	"go.kicksware.com/api/service-common/core/meta"
)

type ReferenceSearchService interface {
	Search(query string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	SearchBy(field, query string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	SearchSKU(sku string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	SearchBrand(brand string, params *meta.RequestParams) ([]*model.SneakerReference, error)
	SearchModel(model string, params *meta.RequestParams) ([]*model.SneakerReference, error)
}