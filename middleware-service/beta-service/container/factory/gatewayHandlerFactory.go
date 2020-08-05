package factory

import
(
	"github.com/go-chi/chi"

	grpc "beta-service/api/gRPC"
	"beta-service/api/rest"
	"beta-service/core/service"
	"beta-service/env"
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