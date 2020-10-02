package factory

import (
	"go.kicksware.com/api/service-common/api/rest"
	"go.kicksware.com/api/service-common/core"

	"go.kicksware.com/api/product-service/core/repo"
	"go.kicksware.com/api/product-service/core/service"
	"go.kicksware.com/api/product-service/env"
	"go.kicksware.com/api/product-service/usecase/business"
)

func ProvideDataService(repository repo.SneakerProductRepository, auth core.AuthService, config env.ServiceConfig) service.SneakerProductService {
	return business.NewSneakerProductService(repository, auth, config.Common)
}

func ProvideAuthService(config env.ServiceConfig) core.AuthService {
	return rest.NewAuthService(config.Auth)
}