package factory

import (
	"github.com/go-chi/chi"

	"search-service/api/rest"
	"search-service/core/service"
	"search-service/env"
)

func ProvideGatewayHandler(search service.ReferenceSearchService, sync service.ReferenceSyncService, auth service.AuthService, config env.ServiceConfig) rest.RestfulHandler {
	return rest.NewHandler(search, sync, auth, config.Common)
}

func ProvideEndpointRouter(handler rest.RestfulHandler) chi.Router {
	return rest.ProvideRoutes(handler)
}