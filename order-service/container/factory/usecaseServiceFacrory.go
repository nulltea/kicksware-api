package factory

import (
	"go.kicksware.com/api/search-service/core/pipe"
	"go.kicksware.com/api/service-common/api/rest"
	"go.kicksware.com/api/service-common/core"

	"go.kicksware.com/api/order-service/core/repo"
	"go.kicksware.com/api/order-service/core/service"
	"go.kicksware.com/api/order-service/env"
	"go.kicksware.com/api/order-service/usecase/business"
)

func ProvideDataService(repository repo.OrderRepository, pipe pipe.SneakerReferencePipe, auth core.AuthService, config env.ServiceConfig) service.OrderService {
	return business.NewOrderService(repository, pipe, auth, config)
}

func ProvideAuthService(config env.ServiceConfig) core.AuthService {
	return rest.NewAuthService(config.Auth)
}