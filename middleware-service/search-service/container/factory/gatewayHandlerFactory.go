package factory

import (
	"github.com/go-chi/chi"

	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/api/REST"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/api/gRPC"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/service"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/env"
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