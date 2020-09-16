package config

import (
	"github.com/timoth-y/kicksware-api/service-common/container"

	"github.com/timoth-y/kicksware-api/user-service/container/factory"
	"github.com/timoth-y/kicksware-api/user-service/env"
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
