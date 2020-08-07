package factory

import (
	"github.com/go-chi/chi"

	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/core"
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/server"

	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/env"
)

func ProvideServer(config env.ServiceConfig, router chi.Router) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupRoutes(router)
	return srv
}
