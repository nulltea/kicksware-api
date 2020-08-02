package factory

import (
	"github.com/go-chi/chi"

	"reference-service/api/gRPC"
	"reference-service/api/rest"
	"reference-service/core/service"
	"reference-service/env"
)

func ProvideRESTGatewayHandler(service service.SneakerReferenceService, auth service.AuthService, config env.ServiceConfig) *rest.Handler {
	return rest.NewHandler(service, auth, config.Common)
}

func ProvideGRPCGatewayHandler(service service.SneakerReferenceService, auth service.AuthService, config env.ServiceConfig) *gRPC.Handler {
	return gRPC.NewHandler(service, auth, config.Common)

}

func ProvideEndpointRouter(handler *rest.Handler) chi.Router {
	return rest.ProvideRoutes(handler)
}