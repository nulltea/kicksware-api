package service

import "go.kicksware.com/api/services/products/core/model"

type ProductSearchService interface {
	Search(query string) ([]*model.SneakerProduct, error)
	SearchBy(field, query string) ([]*model.SneakerProduct, error)
	SearchSKU(sku string) ([]*model.SneakerProduct, error)
	SearchBrand(brand string) ([]*model.SneakerProduct, error)
	SearchModel(model string) ([]*model.SneakerProduct, error)
}
