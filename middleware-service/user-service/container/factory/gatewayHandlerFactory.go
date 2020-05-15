package factory

import (
	"github.com/go-chi/chi"

	"user-service/api/rest"
	"user-service/core/service"
	"user-service/env"
)

func ProvideGatewayHandler(service service.UserService, authService service.AuthService, mailService service.MailService,
	config env.ServiceConfig) rest.RestfulHandler {
	return rest.NewHandler(
		service,
		authService,
		mailService,
		config.Common,
	)
}

func ProvideEndpointRouter(handler rest.RestfulHandler) chi.Router {
	return rest.ProvideRoutes(handler)
}