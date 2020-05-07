package factory

import (
	"product-service/core/repo"
	"product-service/core/service"
	"product-service/env"
	"product-service/usecase/business"
)

func ProvideDataService(repository repo.SneakerProductRepository, config env.ServiceConfig) service.SneakerProductService {
	return business.NewSneakerProductService(repository, config.Common)
}

func ProvideAuthService(config env.ServiceConfig) service.AuthService {
	return business.NewAuthService(config.Auth)
}