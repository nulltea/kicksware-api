package config

import (
	"go.kicksware.com/api/shared/container"

	"go.kicksware.com/api/services/users/container/factory"
	"go.kicksware.com/api/services/users/env"
)

func ConfigureContainer(container container.ServiceContainer, config env.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(factory.ProvideRepository).
		BindSingleton(factory.ProvideLikesRepository).
		BindSingleton(factory.ProvideRemotesRepository).
		BindSingleton(factory.ProvideSubscriptionsRepository).

		BindSingleton(factory.ProvideDataService).
		BindSingleton(factory.ProvideAuthService).
		BindSingleton(factory.ProvideMailService).
		BindSingleton(factory.ProvideInteractService).

		BindSingleton(factory.ProvideRESTGatewayHandler).
		BindSingleton(factory.ProvideGRPCGatewayHandler).
		BindTransient(factory.ProvideEndpointRouter).

		BindTransient(factory.ProvideServer)
}
