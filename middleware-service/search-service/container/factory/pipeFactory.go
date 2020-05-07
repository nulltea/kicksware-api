package factory

import (
	"search-service/core/pipe"
	"search-service/env"
	"search-service/usecase/pipes"
)

func ProvideReferencePipe(config env.ServiceConfig) pipe.SneakerReferencePipe {
	return pipes.NewSneakerReferencePipe(config.Common)
}

func ProvideProductPipe(config env.ServiceConfig) pipe.SneakerProductPipe {
	// return pipes.NewSneakerProductPipe(config.Common)
	panic("ProvideProductPipe: not implemented")
}

