package config

import (
	"go.kicksware.com/api/shared/container"

	"go.kicksware.com/api/services/references/container/factory"
	"go.kicksware.com/api/services/references/env"
)

func ConfigureContainer(container container.ServiceContainer, config env.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(factory.ProvideRepository).

		BindSingleton(factory.ProvideAuthService).
		BindSingleton(factory.ProvideDataService).

		BindSingleton(factory.ProvideRESTGatewayHandler).
		BindSingleton(factory.ProvideGRPCGatewayHandler).

		BindTransient(factory.ProvideEndpointRouter).

		BindTransient(factory.ProvideServer)
}
