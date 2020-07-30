package factory

import (
	"github.com/go-chi/chi"
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/core"
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/server"

	"user-service/env"
)

func ProvideServer(config env.ServiceConfig, router chi.Router) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupREST(router)
	// srv.SetupGRPC(grpc.ProvideRemoteSetup(gRpc))
	return srv
}