package service

import "search-service/core/model"

type ReferenceSearchService interface {
	Search(query string) ([]*model.SneakerReference, error)
	SearchBy(field, query string) ([]*model.SneakerReference, error)
	SearchSKU(sku string) ([]*model.SneakerReference, error)
	SearchBrand(brand string) ([]*model.SneakerReference, error)
	SearchModel(model string) ([]*model.SneakerReference, error)
}