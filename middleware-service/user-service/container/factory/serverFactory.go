package factory

import (
	"github.com/go-chi/chi"
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/core"
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/server"

	grpc "user-service/api/gRPC"
	"user-service/env"
)

func ProvideServer(config env.ServiceConfig, router chi.Router, gRpc *grpc.Handler) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupREST(router)
	srv.SetupGRPC(grpc.ProvideRemoteSetup(gRpc))
	return srv
}