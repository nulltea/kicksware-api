package config

import (
	"go.kicksware.com/api/shared/container"

	"go.kicksware.com/api/services/rating/container/factory"
	"go.kicksware.com/api/services/rating/env"
)

func ConfigureContainer(container container.ServiceContainer, config env.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(factory.ProvideRepository).

		BindSingleton(factory.ProvideDataService).
		BindSingleton(factory.ProvideAuthService).

		BindSingleton(factory.ProvideEventBus)
}
