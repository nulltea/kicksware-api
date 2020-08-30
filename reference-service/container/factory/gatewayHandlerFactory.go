package factory

import (
	"github.com/go-chi/chi"

	"github.com/timoth-y/kicksware-platform/middleware-service/reference-service/api/gRPC"
	"github.com/timoth-y/kicksware-platform/middleware-service/reference-service/api/rest"
	"github.com/timoth-y/kicksware-platform/middleware-service/reference-service/core/service"
	"github.com/timoth-y/kicksware-platform/middleware-service/reference-service/env"
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