package factory

import (
	"go.kicksware.com/api/shared/api/rest"
	"go.kicksware.com/api/shared/core"

	"go.kicksware.com/api/services/products/core/repo"
	"go.kicksware.com/api/services/products/core/service"
	"go.kicksware.com/api/services/products/env"
	"go.kicksware.com/api/services/products/usecase/business"
)

func ProvideDataService(repository repo.SneakerProductRepository, auth core.AuthService, config env.ServiceConfig) service.SneakerProductService {
	return business.NewSneakerProductService(repository, auth, config.Common)
}

func ProvideAuthService(config env.ServiceConfig) core.AuthService {
	return rest.NewAuthService(config.Auth)
}
