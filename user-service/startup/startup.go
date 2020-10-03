package startup

import (
	di "go.kicksware.com/api/service-common/container"
	"go.kicksware.com/api/service-common/core"

	conf "go.kicksware.com/api/user-service/container/config"
	"go.kicksware.com/api/user-service/env"
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
