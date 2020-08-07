package factory

import (
	"github.com/go-chi/chi"
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/core"
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/server"

	grpc "github.com/timoth-y/kicksware-platform/middleware-service/user-service/api/gRPC"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/env"
)

func ProvideServer(config env.ServiceConfig, router chi.Router, handler *grpc.Handler) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupAuth(handler.ProvideAuthKey(), handler.ProvideAccessRoles())
	srv.SetupREST(router)
	srv.SetupGRPC(grpc.ProvideRemoteSetup(handler))
	return srv
}