package factory

import (
	"github.com/go-chi/chi"

	grpc "github.com/timoth-y/kicksware-api/beta-service/api/gRPC"
	"github.com/timoth-y/kicksware-api/beta-service/api/rest"
	"github.com/timoth-y/kicksware-api/beta-service/core/service"
	"github.com/timoth-y/kicksware-api/beta-service/env"
)

func ProvideRESTGatewayHandler(service service.BetaService, auth service.AuthService, config env.ServiceConfig) *rest.Handler {
	return rest.NewHandler(service, auth, config.Common)
}

func ProvideGRPCGatewayHandler(service service.BetaService, auth service.AuthService, config env.ServiceConfig) *grpc.Handler {
	return grpc.NewHandler(
		service,
		auth,
		config.Common,
	)
}

func ProvideEndpointRouter(handler *rest.Handler) chi.Router {
	return rest.ProvideRoutes(handler)
}