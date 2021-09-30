package config

import (
	"go.kicksware.com/api/shared/container"

	"go.kicksware.com/api/services/search/container/factory"
	"go.kicksware.com/api/services/search/env"
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
