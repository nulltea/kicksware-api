package service

import "elastic-search-service/core/model"

type SneakerReferenceService interface {
	SyncOne(code string) error
	Sync(codes []string) error
	SyncAll() error
	SyncQuery(query interface{}) error
	Search(query string) ([]*model.SneakerReference, error)
	SearchBy(field, query string) ([]*model.SneakerReference, error)
	SearchSKU(sku string) ([]*model.SneakerReference, error)
	SearchBrand(brand string) ([]*model.SneakerReference, error)
	SearchModel(model string) ([]*model.SneakerReference, error)
}