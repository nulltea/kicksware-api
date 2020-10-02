package startup

import (
	di "github.com/timoth-y/kicksware-api/service-common/container"
	"github.com/timoth-y/kicksware-api/service-common/core"

	conf "go.kicksware.com/api/beta-service/container/config"
	"go.kicksware.com/api/beta-service/env"
)

func InitializeServer() (srv core.Server) {
	env.InitEnvironment()
	config, err := env.ReadServiceConfig(env.ServiceConfigPath); if err != nil {
		return nil
	}
	container := di.NewServiceContainer()
	conf.ConfigureContainer(container, config)
	container.Resolve(&srv)
	return
}

