package factory

import (
	"github.com/go-chi/chi"

	"user-service/api/rest"
	"user-service/core/service"
	"user-service/env"
)

func ProvideGatewayHandler(service service.UserService, authService service.AuthService,
	config env.ServiceConfig) rest.RestfulHandler {
	return rest.NewHandler(service, authService, config.Common.ContentType)
}

func ProvideEndpointRouter(handler rest.RestfulHandler) chi.Router {
	return rest.ProvideRoutes(handler)
}