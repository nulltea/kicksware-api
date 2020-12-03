package config

import (
	"go.kicksware.com/api/service-common/container"

	"go.kicksware.com/api/rating-service/container/factory"
	"go.kicksware.com/api/rating-service/env"
)

func ConfigureContainer(container container.ServiceContainer, config env.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(factory.ProvideRepository).

		BindSingleton(factory.ProvideDataService).
		BindSingleton(factory.ProvideAuthService).

		BindSingleton(factory.ProvideEventBus)
}
