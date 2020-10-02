package config

import (
	"go.kicksware.com/api/service-common/container"

	"go.kicksware.com/api/cdn-service/container/factory"
	"go.kicksware.com/api/cdn-service/env"
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
