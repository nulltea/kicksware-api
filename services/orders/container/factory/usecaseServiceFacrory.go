package factory

import (
	"go.kicksware.com/api/services/search/core/pipe"
	"go.kicksware.com/api/shared/api/rest"
	"go.kicksware.com/api/shared/core"

	"go.kicksware.com/api/services/orders/core/repo"
	"go.kicksware.com/api/services/orders/core/service"
	"go.kicksware.com/api/services/orders/env"
	"go.kicksware.com/api/services/orders/usecase/business"
)

func ProvideDataService(repository repo.OrderRepository, pipe pipe.SneakerReferencePipe, auth core.AuthService, config env.ServiceConfig) service.OrderService {
	return business.NewOrderService(repository, pipe, auth, config)
}

func ProvideAuthService(config env.ServiceConfig) core.AuthService {
	return rest.NewAuthService(config.Auth)
}
