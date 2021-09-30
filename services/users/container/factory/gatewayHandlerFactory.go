package factory

import (
	"github.com/go-chi/chi"

	grpc "go.kicksware.com/api/services/users/api/gRPC"
	"go.kicksware.com/api/services/users/api/rest"
	"go.kicksware.com/api/services/users/core/service"
	"go.kicksware.com/api/services/users/env"
)

func ProvideRESTGatewayHandler(service service.UserService, authService service.AuthService, mailService service.MailService,
	interactService service.InteractService, config env.ServiceConfig) *rest.Handler {
	return rest.NewHandler(
		service,
		authService,
		mailService,
		interactService,
		config.Common,
	)
}

func ProvideGRPCGatewayHandler(service service.UserService, authService service.AuthService,
mailService service.MailService, interactService service.InteractService) *grpc.Handler {
	return grpc.NewHandler(
		service,
		authService,
		mailService,
		interactService,
	)
}

func ProvideEndpointRouter(handler *rest.Handler) chi.Router {
	return rest.ProvideRoutes(handler)
}
