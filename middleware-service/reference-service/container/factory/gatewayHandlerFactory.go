package factory

import (
	"github.com/go-chi/chi"

	"reference-service/api/rest"
	"reference-service/core/service"
	"reference-service/env"
)

func ProvideGatewayHandler(service service.SneakerReferenceService, auth service.AuthService, config env.ServiceConfig) rest.RestfulHandler {
	return rest.NewHandler(service, auth, config.Common)
}

func ProvideEndpointRouter(handler rest.RestfulHandler) chi.Router {
	return rest.ProvideRoutes(handler)
}