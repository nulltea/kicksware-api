package factory

import (
	"github.com/go-chi/chi"

	"product-service/api/gRPC"
	"product-service/api/rest"
	"product-service/core/service"
	"product-service/env"
)

func ProvideRESTGatewayHandler(service service.SneakerProductService, auth service.AuthService, config env.ServiceConfig) *rest.Handler {
	return rest.NewHandler(service, auth, config.Common)
}

func ProvideGRPCGatewayHandler(service service.SneakerProductService, auth service.AuthService, config env.ServiceConfig) *gRPC.Handler {
	return gRPC.NewHandler(service, auth, config.Common)
}

func ProvideEndpointRouter(handler *rest.Handler) chi.Router {
	return rest.ProvideRoutes(handler)
}