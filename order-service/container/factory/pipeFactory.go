package factory

import (
	"encoding/json"

	"github.com/timoth-y/kicksware-api/search-service/core/pipe"
	"github.com/timoth-y/kicksware-api/search-service/core/service"
	searcnEnv "github.com/timoth-y/kicksware-api/search-service/env"
	"github.com/timoth-y/kicksware-api/search-service/usecase/pipes/REST"
	"github.com/timoth-y/kicksware-api/search-service/usecase/pipes/gRPC"

	"github.com/timoth-y/kicksware-api/order-service/env"
)

func ProvideReferenceRESTPipe(auth service.AuthService, config env.ServiceConfig) pipe.SneakerReferencePipe {
	var searchConfig searcnEnv.ServiceConfig
	castService(config, &searchConfig)
	return REST.NewSneakerReferencePipe(auth, searchConfig.Common)
}

func ProvideReferenceGRPCPipe(config env.ServiceConfig) pipe.SneakerReferencePipe {
	var searchConfig searcnEnv.ServiceConfig
	castService(config, &searchConfig)
	return gRPC.NewSneakerReferencePipe(searchConfig)
}

func ProvideProductRESTPipe(config env.ServiceConfig) pipe.SneakerProductPipe {
	// return pipes.NewSneakerProductPipe(config.Common)
	panic("ProvideProductRESTPipe: not implemented")
}

func ProvideProductGRPCPipe(config env.ServiceConfig) pipe.SneakerProductPipe {
	// return pipes.NewSneakerProductPipe(config.Common)
	panic("ProvideProductGRPCPipe: not implemented")
}

func castService(native interface{}, foreign interface{}) error {
	bytes, err := json.Marshal(native); if err != nil {
		return err
	}
	json.Unmarshal(bytes, foreign)
	return nil
}
