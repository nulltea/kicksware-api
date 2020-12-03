package factory

import (
	"go.kicksware.com/api/service-common/core"
	"go.kicksware.com/api/service-common/server"

	"go.kicksware.com/api/rating-service/env"
)

func ProvideServer(config env.ServiceConfig) core.Server {
	srv := server.NewInstance(config.Common.Host)
	return srv
}
