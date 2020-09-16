package config

import (
	"github.com/timoth-y/kicksware-api/service-common/container"

	"github.com/timoth-y/kicksware-api/cdn-service/container/factory"
	"github.com/timoth-y/kicksware-api/cdn-service/env"
)

func ConfigureContainer(container container.ServiceContainer, config env.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(factory.ProvideRepository).
		BindSingleton(factory.ProvideContentService).
		BindSingleton(factory.ProvideAuthService).
		BindSingleton(factory.ProvideGatewayHandler).
		BindTransient(factory.ProvideEndpointRouter).
		BindTransient(factory.ProvideServer)
}
