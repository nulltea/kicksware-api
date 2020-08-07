package factory

import (
	"github.com/go-chi/chi"

	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/core"
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/server"

	"github.com/timoth-y/kicksware-platform/middleware-service/beta-service/api/gRPC"

	"github.com/timoth-y/kicksware-platform/middleware-service/beta-service/env"
)

func ProvideServer(config env.ServiceConfig, router chi.Router, handler *grpc.Handler) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupREST(router)
	srv.SetupGRPC(grpc.ProvideRemoteSetup(handler))
	return srv
}
