package startup

import (
	"github.com/golang/glog"
	di "github.com/timoth-y/kicksware-platform/middleware-service/service-common/container"
	"github.com/timoth-y/kicksware-platform/middleware-service/service-common/core"

	conf "github.com/timoth-y/kicksware-platform/middleware-service/user-service/container/config"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/env"
)

func InitializeServer() (srv core.Server) {
	env.InitEnvironment()
	glog.Warningln(env.ServiceConfigPath)
	config, err := env.ReadServiceConfig(env.ServiceConfigPath); if err != nil {
		return nil
	}
	container := di.NewServiceContainer()
	conf.ConfigureContainer(container, config)
	container.Resolve(&srv)
	return
}

