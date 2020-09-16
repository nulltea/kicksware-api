package factory

import (
	"github.com/go-chi/chi"

	"github.com/timoth-y/kicksware-api/service-common/core"
	"github.com/timoth-y/kicksware-api/service-common/server"

	"github.com/timoth-y/kicksware-api/product-service/api/gRPC"
	"github.com/timoth-y/kicksware-api/product-service/env"
)

func ProvideServer(config env.ServiceConfig, router chi.Router, handler *gRPC.Handler) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupEncryption(config.Security.TLSCertificate)
	srv.SetupAuth(handler.ProvideAuthKey(), handler.ProvideAccessRoles())
	srv.SetupREST(router)
	srv.SetupGRPC(gRPC.ProvideRemoteSetup(handler))
	return srv
}
