package factory

import (
	"encoding/json"

	"go.kicksware.com/api/services/search/core/pipe"
	searchEnv "go.kicksware.com/api/services/search/env"
	"go.kicksware.com/api/services/search/usecase/pipes/REST"
	"go.kicksware.com/api/services/search/usecase/pipes/gRPC"
	"go.kicksware.com/api/shared/core"

	"go.kicksware.com/api/services/orders/env"
)

func ProvideReferenceRESTPipe(auth core.AuthService, config env.ServiceConfig) pipe.SneakerReferencePipe {
	return REST.NewSneakerReferencePipe(auth, config.Common)
}

func ProvideReferenceGRPCPipe(config env.ServiceConfig, service core.AuthService) pipe.SneakerReferencePipe {
	var searchConfig searchEnv.ServiceConfig; castService(config, &searchConfig)
	return gRPC.NewSneakerReferencePipe(searchConfig, service)
}

func ProvideProductRESTPipe(config env.ServiceConfig) pipe.SneakerProductPipe {
	// return pipes.NewSneakerProductPipe(config.Common)
	panic("ProvideProductRESTPipe: not implemented")
}

func ProvideProductGRPCPipe(config env.ServiceConfig, service core.AuthService) pipe.SneakerProductPipe {
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
