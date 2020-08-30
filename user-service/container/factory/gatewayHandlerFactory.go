package factory

import (
	"github.com/go-chi/chi"

	grpc "github.com/timoth-y/kicksware-platform/middleware-service/user-service/api/gRPC"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/api/rest"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/service"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/env"
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

func ProvideGRPCGatewayHandler(service service.UserService, authService service.AuthService, mailService service.MailService,
	interactService service.InteractService, config env.ServiceConfig) *grpc.Handler {
	return grpc.NewHandler(
		service,
		authService,
		mailService,
		interactService,
		config.Common,
	)
}

func ProvideEndpointRouter(handler *rest.Handler) chi.Router {
	return rest.ProvideRoutes(handler)
}