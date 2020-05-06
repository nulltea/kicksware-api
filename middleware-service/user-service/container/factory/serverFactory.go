package factory

import (
	"github.com/go-chi/chi"

	"user-service/core/service"
	"user-service/env"
	"user-service/server"
)

func ProvideServer(config env.ServiceConfig, router chi.Router) service.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupRoutes(router)
	return srv
}