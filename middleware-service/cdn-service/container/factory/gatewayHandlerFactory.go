package factory

import (
	"github.com/go-chi/chi"

	"cdn-service/api/rest"
	"cdn-service/core/service"
	"cdn-service/env"
)

func ProvideGatewayHandler(service service.ContentService, auth service.AuthService, config env.ServiceConfig) rest.RestfulHandler {
	return rest.NewHandler(service, auth, config.Common)
}

func ProvideEndpointRouter(handler rest.RestfulHandler) chi.Router {
	return rest.ProvideRoutes(handler)
}