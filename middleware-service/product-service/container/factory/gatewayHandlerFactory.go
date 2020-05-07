package factory

import (
	"github.com/go-chi/chi"

	"product-service/api/rest"
	"product-service/core/service"
	"product-service/env"
)

func ProvideGatewayHandler(service service.SneakerProductService, auth service.AuthService, config env.ServiceConfig) rest.RestfulHandler {
	return rest.NewHandler(service, auth, config.Common)
}

func ProvideEndpointRouter(handler rest.RestfulHandler) chi.Router {
	return rest.ProvideRoutes(handler)
}