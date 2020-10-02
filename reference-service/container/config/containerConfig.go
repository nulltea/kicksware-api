package config

import (
	"go.kicksware.com/api/service-common/container"

	"go.kicksware.com/api/reference-service/container/factory"
	"go.kicksware.com/api/reference-service/env"
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
