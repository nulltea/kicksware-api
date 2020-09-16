package factory

import (
	"github.com/go-chi/chi"

	"github.com/timoth-y/kicksware-api/service-common/core"
	"github.com/timoth-y/kicksware-api/service-common/server"

	"github.com/timoth-y/kicksware-api/cdn-service/env"
)

func ProvideServer(config env.ServiceConfig, router chi.Router) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupREST(router)
	return srv
}
