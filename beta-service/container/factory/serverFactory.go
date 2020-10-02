package factory

import (
	"github.com/go-chi/chi"
	"github.com/timoth-y/kicksware-api/service-common/core"
	"github.com/timoth-y/kicksware-api/service-common/server"

	"go.kicksware.com/api/beta/api/gRPC"

	"go.kicksware.com/api/beta/env"
)

func ProvideServer(config env.ServiceConfig, router chi.Router, handler *grpc.Handler) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupAuth(handler.ProvideAuthKey(), handler.ProvideAccessRoles())
	srv.SetupREST(router)
	srv.SetupGRPC(grpc.ProvideRemoteSetup(handler))
	return srv
}
