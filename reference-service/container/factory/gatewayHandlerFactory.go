package factory

import (
	"github.com/go-chi/chi"
	"go.kicksware.com/api/service-common/core"

	"go.kicksware.com/api/reference-service/api/gRPC"
	"go.kicksware.com/api/reference-service/api/rest"
	"go.kicksware.com/api/reference-service/core/service"
	"go.kicksware.com/api/reference-service/env"
)

func ProvideRESTGatewayHandler(service service.SneakerReferenceService, auth core.AuthService, config env.ServiceConfig) *rest.Handler {
	return rest.NewHandler(service, auth, config.Common)
}

func ProvideGRPCGatewayHandler(service service.SneakerReferenceService, auth core.AuthService, config env.ServiceConfig) *gRPC.Handler {
	return gRPC.NewHandler(service, auth, config.Common)

}

func ProvideEndpointRouter(handler *rest.Handler) chi.Router {
	return rest.ProvideRoutes(handler)
}