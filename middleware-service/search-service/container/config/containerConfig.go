package config

import (
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/container"

	"search-service/container/factory"
	"search-service/env"
)

func ConfigureContainer(container container.ServiceContainer, config env.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(factory.ProvideReferencePipe).
		BindSingleton(factory.ProvideProductPipe).

		BindSingleton(factory.ProvideReferenceSearchService).
		BindSingleton(factory.ProvideProductSearchService).

		BindSingleton(factory.ProvideReferenceSyncService).
		BindSingleton(factory.ProvideProductSyncService).

		BindSingleton(factory.ProvideAuthService).

		BindSingleton(factory.ProvideGatewayHandler).
		BindTransient(factory.ProvideEndpointRouter).

		BindTransient(factory.ProvideServer)
}
