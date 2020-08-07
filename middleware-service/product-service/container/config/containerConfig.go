package config

import (
	"github.com/timoth-y/kicksware-platform/middleware-service/service-common/container"

	"github.com/timoth-y/kicksware-platform/middleware-service/product-service/container/factory"
	"github.com/timoth-y/kicksware-platform/middleware-service/product-service/env"
)

func ConfigureContainer(container container.ServiceContainer, config env.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(factory.ProvideRepository).

		BindSingleton(factory.ProvideDataService).
		BindSingleton(factory.ProvideAuthService).

		BindSingleton(factory.ProvideRESTGatewayHandler).
		BindTransient(factory.ProvideEndpointRouter).
		BindSingleton(factory.ProvideGRPCGatewayHandler).

		BindTransient(factory.ProvideServer)
}
