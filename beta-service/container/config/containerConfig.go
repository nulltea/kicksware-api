package config

import (
	"github.com/timoth-y/kicksware-api/service-common/container"

	"go.kicksware.com/api/beta/container/factory"
	"go.kicksware.com/api/beta/env"
)

func ConfigureContainer(container container.ServiceContainer, config env.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(factory.ProvideRepository).

		BindSingleton(factory.ProvideDataService).
		BindSingleton(factory.ProvideAuthService).

		BindSingleton(factory.ProvideRESTGatewayHandler).
		BindSingleton(factory.ProvideGRPCGatewayHandler).

		BindTransient(factory.ProvideEndpointRouter).

		BindTransient(factory.ProvideServer)
}
