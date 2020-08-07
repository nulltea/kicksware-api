package factory

import (
	"github.com/timoth-y/kicksware-platform/middleware-service/product-service/core/repo"
	"github.com/timoth-y/kicksware-platform/middleware-service/product-service/core/service"
	"github.com/timoth-y/kicksware-platform/middleware-service/product-service/env"
	"github.com/timoth-y/kicksware-platform/middleware-service/product-service/usecase/business"
)

func ProvideDataService(repository repo.SneakerProductRepository, config env.ServiceConfig) service.SneakerProductService {
	return business.NewSneakerProductService(repository, config.Common)
}

func ProvideAuthService(config env.ServiceConfig) service.AuthService {
	return business.NewAuthService(config.Auth)
}