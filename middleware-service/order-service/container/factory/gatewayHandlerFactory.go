package factory

import (
	"github.com/go-chi/chi"

	"order-service/api/rest"
	"order-service/core/service"
	"order-service/env"
)

func ProvideGatewayHandler(service service.OrderService, auth service.AuthService, config env.ServiceConfig) rest.RestfulHandler {
	return rest.NewHandler(service, auth, config.Common)
}

func ProvideEndpointRouter(handler rest.RestfulHandler) chi.Router {
	return rest.ProvideRoutes(handler)
}