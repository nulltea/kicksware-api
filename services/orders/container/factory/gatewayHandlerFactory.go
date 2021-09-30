package factory

import (
	"github.com/go-chi/chi"
	"go.kicksware.com/api/shared/core"

	"go.kicksware.com/api/services/orders/api/gRPC"
	"go.kicksware.com/api/services/orders/api/rest"
	"go.kicksware.com/api/services/orders/core/service"
	"go.kicksware.com/api/services/orders/env"
)

func ProvideRESTGatewayHandler(service service.OrderService, auth core.AuthService, config env.ServiceConfig) *rest.Handler {
	return rest.NewHandler(service, auth, config.Common)
}

func ProvideGRPCGatewayHandler(service service.OrderService, auth core.AuthService) *gRPC.Handler {
	return gRPC.NewHandler(service, auth)
}

func ProvideEndpointRouter(handler *rest.Handler) chi.Router {
	return rest.ProvideRoutes(handler)
}
