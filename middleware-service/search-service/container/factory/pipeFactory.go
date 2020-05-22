package factory

import (
	"search-service/core/pipe"
	"search-service/core/service"
	"search-service/env"
	"search-service/usecase/pipes"
)

func ProvideReferencePipe(auth service.AuthService, config env.ServiceConfig) pipe.SneakerReferencePipe {
	return pipes.NewSneakerReferencePipe(auth, config.Common)
}

func ProvideProductPipe(config env.ServiceConfig) pipe.SneakerProductPipe {
	// return pipes.NewSneakerProductPipe(config.Common)
	panic("ProvideProductPipe: not implemented")
}

