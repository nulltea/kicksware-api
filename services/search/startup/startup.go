package startup

import (
	"log"

	"github.com/pkg/errors"
	di "go.kicksware.com/api/shared/container"
	"go.kicksware.com/api/shared/core"

	conf "go.kicksware.com/api/services/search/container/config"
	"go.kicksware.com/api/services/search/core/service"
	"go.kicksware.com/api/services/search/env"
)

func InitializeServer() (srv core.Server, container di.ServiceContainer) {
	env.InitEnvironment()
	config, err := env.ReadServiceConfig(env.ServiceConfigPath); if err != nil {
		return
	}
	container = di.NewServiceContainer()
	conf.ConfigureContainer(container, config)
	container.Resolve(&srv)
	return
}

func PerformDataSync(container di.ServiceContainer) error {
	return container.ResolveFor(func (service service.ReferenceSyncService) error {
		if err := service.SyncAll(nil); err != nil {
			log.Println(errors.Wrap(err, "search-service::startup.PerformDataSync: sneaker references replication sync failed"))
			return errors.Wrap(err, "search-service::startup.PerformDataSync: sneaker references replication sync failed")
		}
		log.Println("search-service::startup.PerformDataSync: sneaker references replication sync completed with success")
		return nil
	})
}
