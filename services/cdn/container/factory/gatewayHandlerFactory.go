package factory

import (
	"github.com/go-chi/chi"
	"go.kicksware.com/api/shared/core"

	"go.kicksware.com/api/services/cdn/api/gRPC"
	"go.kicksware.com/api/services/cdn/api/rest"
	"go.kicksware.com/api/services/cdn/core/service"
)

func ProvideGatewayHandler(service service.ContentService, auth core.AuthService) *rest.Handler {
	return rest.NewHandler(service, auth)
}

func ProvideGRPCGatewayHandler(service service.ContentService, auth core.AuthService) *gRPC.Handler {
	return gRPC.NewHandler(service, auth)
}

func ProvideEndpointRouter(handler *rest.Handler) chi.Router {
	return rest.ProvideRoutes(handler)
}
