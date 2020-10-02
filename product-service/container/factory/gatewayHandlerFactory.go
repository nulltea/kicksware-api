package factory

import (
	"github.com/go-chi/chi"
	"go.kicksware.com/api/service-common/core"

	"go.kicksware.com/api/product-service/api/gRPC"
	"go.kicksware.com/api/product-service/api/rest"
	"go.kicksware.com/api/product-service/core/service"
	"go.kicksware.com/api/product-service/env"
)

func ProvideRESTGatewayHandler(service service.SneakerProductService, auth core.AuthService, config env.ServiceConfig) *rest.Handler {
	return rest.NewHandler(service, auth, config.Common)
}

func ProvideGRPCGatewayHandler(service service.SneakerProductService, auth core.AuthService, config env.ServiceConfig) *gRPC.Handler {
	return gRPC.NewHandler(service, auth, config.Common)
}

func ProvideEndpointRouter(handler *rest.Handler) chi.Router {
	return rest.ProvideRoutes(handler)
}