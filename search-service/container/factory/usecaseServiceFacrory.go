package factory

import (
	"go.kicksware.com/api/service-common/api/rest"
	"go.kicksware.com/api/service-common/core"

	"go.kicksware.com/api/search-service/core/pipe"
	"go.kicksware.com/api/search-service/core/service"
	"go.kicksware.com/api/search-service/env"
	"go.kicksware.com/api/search-service/usecase/business"
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

func ProvideAuthService(config env.ServiceConfig) core.AuthService {
	return rest.NewAuthService(config.Auth)
}