package factory

import (
	"order-service/core/repo"
	"order-service/core/service"
	"order-service/env"
	"order-service/usecase/business"
)

func ProvideDataService(repository repo.OrderRepository, config env.ServiceConfig) service.OrderService {
	return business.NewOrderService(repository, config.Common)
}

func ProvideAuthService(config env.ServiceConfig) service.AuthService {
	return business.NewAuthService(config.Auth)
}