package factory

import (
	"encoding/json"

	"go.kicksware.com/api/search-service/core/pipe"
	searchEnv "go.kicksware.com/api/search-service/env"
	"go.kicksware.com/api/search-service/usecase/pipes/REST"
	"go.kicksware.com/api/search-service/usecase/pipes/gRPC"
	"go.kicksware.com/api/service-common/core"

	"go.kicksware.com/api/order-service/env"
)

func ProvideReferenceRESTPipe(auth core.AuthService, config env.ServiceConfig) pipe.SneakerReferencePipe {
	return REST.NewSneakerReferencePipe(auth, config.Common)
}

func ProvideReferenceGRPCPipe(config env.ServiceConfig) pipe.SneakerReferencePipe {
	var searchConfig searchEnv.ServiceConfig; castService(config, &searchConfig)
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
