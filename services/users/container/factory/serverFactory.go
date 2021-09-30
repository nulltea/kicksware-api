package factory

import (
	"github.com/go-chi/chi"
	"go.kicksware.com/api/shared/core"
	"go.kicksware.com/api/shared/server"

	grpc "go.kicksware.com/api/services/users/api/gRPC"
	"go.kicksware.com/api/services/users/env"
)

func ProvideServer(config env.ServiceConfig, router chi.Router, handler *grpc.Handler) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupEncryption(config.Security.TLSCertificate)
	srv.SetupAuth(handler.ProvideAuthKey(), handler.ProvideAccessRoles())
	srv.SetupREST(router)
	srv.SetupGRPC(grpc.ProvideRemoteSetup(handler))
	return srv
}
