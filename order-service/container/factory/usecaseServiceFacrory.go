package factory

import (
	"github.com/timoth-y/kicksware-api/search-service/core/pipe"

	"github.com/timoth-y/kicksware-api/order-service/core/repo"
	"github.com/timoth-y/kicksware-api/order-service/core/service"
	"github.com/timoth-y/kicksware-api/order-service/env"
	"github.com/timoth-y/kicksware-api/order-service/usecase/business"
)

func ProvideDataService(repository repo.OrderRepository, pipe pipe.SneakerReferencePipe, config env.ServiceConfig) service.OrderService {
	return business.NewOrderService(repository, pipe, config.Common)
}

func ProvideAuthService(config env.ServiceConfig) service.AuthService {
	return business.NewAuthService(config.Auth)
}