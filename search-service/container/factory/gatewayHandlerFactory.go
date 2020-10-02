package factory

import (
	"github.com/go-chi/chi"
	"go.kicksware.com/api/service-common/core"

	"go.kicksware.com/api/search-service/api/REST"
	"go.kicksware.com/api/search-service/api/gRPC"
	"go.kicksware.com/api/search-service/core/service"
	"go.kicksware.com/api/search-service/env"
)

func ProvideRESTGatewayHandler(search service.ReferenceSearchService, sync service.ReferenceSyncService, auth core.AuthService, config env.ServiceConfig) *rest.Handler {
	return rest.NewHandler(search, sync, auth, config.Common)
}

func ProvideGRPCGatewayHandler(search service.ReferenceSearchService, sync service.ReferenceSyncService, auth core.AuthService, config env.ServiceConfig) *gRPC.Handler {
	return gRPC.NewHandler(search, sync, auth)
}

func ProvideEndpointRouter(handler *rest.Handler) chi.Router {
	return rest.ProvideRoutes(handler)
}