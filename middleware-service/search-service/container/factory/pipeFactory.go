package factory

import (
	"search-service/core/pipe"
	"search-service/core/service"
	"search-service/env"
	"search-service/usecase/pipes/REST"
	"search-service/usecase/pipes/gRPC"
)

func ProvideReferenceRESTPipe(auth service.AuthService, config env.ServiceConfig) pipe.SneakerReferencePipe {
	return REST.NewSneakerReferencePipe(auth, config.Common)
}

func ProvideReferenceGRPCPipe(auth service.AuthService, config env.ServiceConfig) pipe.SneakerReferencePipe {
	return gRPC.NewSneakerReferencePipe(auth, config.Common)
}

func ProvideProductRESTPipe(config env.ServiceConfig) pipe.SneakerProductPipe {
	// return pipes.NewSneakerProductPipe(config.Common)
	panic("ProvideProductRESTPipe: not implemented")
}

func ProvideProductGRPCPipe(config env.ServiceConfig) pipe.SneakerProductPipe {
	// return pipes.NewSneakerProductPipe(config.Common)
	panic("ProvideProductGRPCPipe: not implemented")
}

