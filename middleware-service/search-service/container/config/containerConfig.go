package config

import (
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/container"

	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/container/factory"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/env"
)

func ConfigureContainer(container container.ServiceContainer, config env.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(factory.ProvideReferenceGRPCPipe).
		// BindSingleton(factory.ProvideProductGRPCPipe).

		BindSingleton(factory.ProvideReferenceSearchService).
		BindSingleton(factory.ProvideProductSearchService).

		BindSingleton(factory.ProvideReferenceSyncService).
		BindSingleton(factory.ProvideProductSyncService).

		BindSingleton(factory.ProvideAuthService).

		BindSingleton(factory.ProvideRESTGatewayHandler).
		BindTransient(factory.ProvideEndpointRouter).
		BindSingleton(factory.ProvideGRPCGatewayHandler).

		BindTransient(factory.ProvideServer)
}
