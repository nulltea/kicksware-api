package config

import (
	"github.com/timoth-y/kicksware-api/service-common/container"

	"github.com/timoth-y/kicksware-api/product-service/container/factory"
	"github.com/timoth-y/kicksware-api/product-service/env"
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
