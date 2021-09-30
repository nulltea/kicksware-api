package config

import (
	"go.kicksware.com/api/shared/container"

	"go.kicksware.com/api/services/cdn/container/factory"
	"go.kicksware.com/api/services/cdn/env"
)

func ConfigureContainer(container container.ServiceContainer, config env.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(factory.ProvideRepository).

		BindSingleton(factory.ProvideContentService).

		BindSingleton(factory.ProvideAuthService).

		BindSingleton(factory.ProvideGatewayHandler).
		BindSingleton(factory.ProvideGRPCGatewayHandler).

		BindTransient(factory.ProvideEndpointRouter).

		BindTransient(factory.ProvideServer)
}
