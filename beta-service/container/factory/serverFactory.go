package factory

import (
	"github.com/go-chi/chi"

	"github.com/timoth-y/kicksware-api/service-common/core"
	"github.com/timoth-y/kicksware-api/service-common/server"

	"github.com/timoth-y/kicksware-api/beta-service/api/gRPC"

	"github.com/timoth-y/kicksware-api/beta-service/env"
)

func ProvideServer(config env.ServiceConfig, router chi.Router, handler *grpc.Handler) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupAuth(handler.ProvideAuthKey(), handler.ProvideAccessRoles())
	srv.SetupREST(router)
	srv.SetupGRPC(grpc.ProvideRemoteSetup(handler))
	return srv
}
