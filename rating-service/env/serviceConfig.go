package env

import (
	"io/ioutil"

	"github.com/golang/glog"
	"go.kicksware.com/api/service-common/config"
	"gopkg.in/yaml.v2"
)

type ServiceConfig struct {
	Common   config.CommonConfig     `yaml:"commonConfig"`
	EventBus config.ConnectionConfig `yaml:"eventBusConfig"`
	Auth     config.AuthConfig       `yaml:"authConfig"`
	Mongo    config.DataStoreConfig  `yaml:"mongoConfig"`
}


func ReadServiceConfig(filename string) (sc ServiceConfig, err error) {
	file, err := ioutil.ReadFile(filename); if err != nil {
		glog.Fatalln(err)
		return
	}

	err = yaml.Unmarshal(file, &sc); if err != nil {
		glog.Fatalln(err)
		return
	}
	return
}
