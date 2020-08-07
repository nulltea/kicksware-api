package factory

import (
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/core/repo"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/core/service"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/env"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/usecase/business"
)

func ProvideDataService(repository repo.OrderRepository, config env.ServiceConfig) service.OrderService {
	return business.NewOrderService(repository, config.Common)
}

func ProvideAuthService(config env.ServiceConfig) service.AuthService {
	return business.NewAuthService(config.Auth)
}