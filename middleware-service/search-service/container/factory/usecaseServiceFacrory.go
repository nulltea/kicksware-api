package factory

import (
	"search-service/core/pipe"
	"search-service/core/service"
	"search-service/env"
	"search-service/usecase/business"
)

func ProvideReferenceSearchService(config env.ServiceConfig) service.ReferenceSearchService {
	return business.NewReferenceSearchService(config.Elastic, config.Search)
}

func ProvideReferenceSyncService(pipe pipe.SneakerReferencePipe, config env.ServiceConfig) service.ReferenceSyncService {
	return business.NewReferenceSyncService(pipe, config.Elastic)
}

func ProvideProductSearchService(config env.ServiceConfig) service.ProductSearchService {
	// return business.NewProductSearchService(config.Elastic)
	panic("ProvideProductDataService: not implemented")
}

func ProvideProductSyncService(pipe pipe.SneakerProductPipe, config env.ServiceConfig) service.ProductSyncService {
	// return business.NewProductSyncService(pipe, config.Common)
	panic("ProvideProductSyncService: not implemented")
}

func ProvideAuthService(config env.ServiceConfig) service.AuthService {
	return business.NewAuthService(config.Auth)
}