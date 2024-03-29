package factory

import (
	"go.kicksware.com/api/shared/core"

	"go.kicksware.com/api/services/search/core/pipe"
	"go.kicksware.com/api/services/search/env"
	"go.kicksware.com/api/services/search/usecase/pipes/REST"
	"go.kicksware.com/api/services/search/usecase/pipes/gRPC"
)

func ProvideReferenceRESTPipe(auth core.AuthService, config env.ServiceConfig) pipe.SneakerReferencePipe {
	return REST.NewSneakerReferencePipe(auth, config.Common)
}

func ProvideReferenceGRPCPipe(config env.ServiceConfig, service core.AuthService) pipe.SneakerReferencePipe {
	return gRPC.NewSneakerReferencePipe(config, service)
}

func ProvideProductRESTPipe(config env.ServiceConfig) pipe.SneakerProductPipe {
	// return pipes.NewSneakerProductPipe(config.Common)
	panic("ProvideProductRESTPipe: not implemented")
}

func ProvideProductGRPCPipe(config env.ServiceConfig, service core.AuthService) pipe.SneakerProductPipe {
	// return pipes.NewSneakerProductPipe(config.Common)
	panic("ProvideProductGRPCPipe: not implemented")
}

