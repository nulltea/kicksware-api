package factory

import (
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/pipe"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/service"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/env"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/usecase/pipes/REST"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/usecase/pipes/gRPC"
)

func ProvideReferenceRESTPipe(auth service.AuthService, config env.ServiceConfig) pipe.SneakerReferencePipe {
	return REST.NewSneakerReferencePipe(auth, config.Common)
}

func ProvideReferenceGRPCPipe(config env.ServiceConfig) pipe.SneakerReferencePipe {
	return gRPC.NewSneakerReferencePipe(config)
}

func ProvideProductRESTPipe(config env.ServiceConfig) pipe.SneakerProductPipe {
	// return pipes.NewSneakerProductPipe(config.Common)
	panic("ProvideProductRESTPipe: not implemented")
}

func ProvideProductGRPCPipe(config env.ServiceConfig) pipe.SneakerProductPipe {
	// return pipes.NewSneakerProductPipe(config.Common)
	panic("ProvideProductGRPCPipe: not implemented")
}

