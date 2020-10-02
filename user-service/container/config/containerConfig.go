package config

import (
	"go.kicksware.com/api/service-common/container"

	"go.kicksware.com/api/user-service/container/factory"
	"go.kicksware.com/api/user-service/env"
)

func ConfigureContainer(container container.ServiceContainer, config env.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(factory.ProvideRepository).
		BindSingleton(factory.ProvideLikesRepository).
		BindSingleton(factory.ProvideRemotesRepository).

		BindSingleton(factory.ProvideDataService).
		BindSingleton(factory.ProvideAuthService).
		BindSingleton(factory.ProvideMailService).
		BindSingleton(factory.ProvideInteractService).

		BindSingleton(factory.ProvideRESTGatewayHandler).
		BindSingleton(factory.ProvideGRPCGatewayHandler).
		BindTransient(factory.ProvideEndpointRouter).

		BindTransient(factory.ProvideServer)
}
