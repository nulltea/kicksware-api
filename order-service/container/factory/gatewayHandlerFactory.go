package factory

import (
	"github.com/go-chi/chi"

	"github.com/timoth-y/kicksware-api/order-service/api/rest"
	"github.com/timoth-y/kicksware-api/order-service/core/service"
	"github.com/timoth-y/kicksware-api/order-service/env"
)

func ProvideGatewayHandler(service service.OrderService, auth service.AuthService, config env.ServiceConfig) rest.RestfulHandler {
	return rest.NewHandler(service, auth, config.Common)
}

func ProvideEndpointRouter(handler rest.RestfulHandler) chi.Router {
	return rest.ProvideRoutes(handler)
}