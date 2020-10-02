package env

import (
	"io/ioutil"

	"github.com/golang/glog"
	"go.kicksware.com/api/service-common/config"
	"gopkg.in/yaml.v2"
)

type ServiceConfig struct {
	Common   config.CommonConfig    `yaml:"commonConfig"`
	Security config.SecurityConfig  `yaml:"securityConfig"`
	Auth     config.AuthConfig      `yaml:"authConfig"`
	Mongo    config.DataStoreConfig `yaml:"mongoConfig"`
	Postgres config.DataStoreConfig `yaml:"postgresConfig"`
	Redis    config.DataStoreConfig `yaml:"redisConfig"`
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
