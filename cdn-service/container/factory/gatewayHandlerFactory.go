package factory

import (
	"github.com/go-chi/chi"

	"github.com/timoth-y/kicksware-api/cdn-service/api/rest"
	"github.com/timoth-y/kicksware-api/cdn-service/core/service"
	"github.com/timoth-y/kicksware-api/cdn-service/env"
)

func ProvideGatewayHandler(service service.ContentService, auth service.AuthService, config env.ServiceConfig) rest.RestfulHandler {
	return rest.NewHandler(service, auth, config.Common)
}

func ProvideEndpointRouter(handler rest.RestfulHandler) chi.Router {
	return rest.ProvideRoutes(handler)
}