package startup

import (
	di "user-service/container"
	conf "user-service/container/config"
	"user-service/core/service"
	"user-service/env"
)

func InitializeServer() (srv service.Server) {
	env.InitEnvironment()
	config, err := env.ReadServiceConfig(env.ServiceConfigPath); if err != nil {
		return nil
	}
	container := di.NewServiceContainer()
	conf.ConfigureContainer(container, config)
	container.Resolve(&srv)
	return
}

