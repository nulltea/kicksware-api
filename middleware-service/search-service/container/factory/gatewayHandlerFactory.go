package factory

import (
	"github.com/go-chi/chi"

	"search-service/api/REST"
	"search-service/api/gRPC"
	"search-service/core/service"
	"search-service/env"
)

func ProvideRESTGatewayHandler(search service.ReferenceSearchService, sync service.ReferenceSyncService, auth service.AuthService, config env.ServiceConfig) *rest.Handler {
	return rest.NewHandler(search, sync, auth, config.Common)
}

func ProvideGRPCGatewayHandler(search service.ReferenceSearchService, sync service.ReferenceSyncService, auth service.AuthService, config env.ServiceConfig) *gRPC.Handler {
	return gRPC.NewHandler(search, sync, auth, config.Common)
}

func ProvideEndpointRouter(handler *rest.Handler) chi.Router {
	return rest.ProvideRoutes(handler)
}