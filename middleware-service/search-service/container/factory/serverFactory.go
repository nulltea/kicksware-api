package factory

import (
	"github.com/go-chi/chi"

	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/core"
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/server"

	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/api/gRPC"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/env"
)

func ProvideServer(config env.ServiceConfig, router chi.Router, handler *gRPC.Handler) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupREST(router)
	srv.SetupGRPC(gRPC.ProvideRemoteSetup(handler))
	return srv
}
