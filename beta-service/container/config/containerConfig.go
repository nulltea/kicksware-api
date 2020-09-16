package config

import (
	"github.com/timoth-y/kicksware-api/service-common/container"

	"github.com/timoth-y/kicksware-api/beta-service/container/factory"
	"github.com/timoth-y/kicksware-api/beta-service/env"
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
