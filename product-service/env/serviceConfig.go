package env

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"github.com/golang/glog"
	"github.com/timoth-y/kicksware-platform/middleware-service/service-common/core/meta"
)

type ServiceConfig struct {
	Common   CommonConfig    `yaml:"commonConfig"`
	Security SecurityConfig  `yaml:"securityConfig"`
	Auth     AuthConfig      `yaml:"authConfig"`
	Mongo    DataStoreConfig `yaml:"mongoConfig"`
	Postgres DataStoreConfig `yaml:"postgresConfig"`
	Redis    DataStoreConfig `yaml:"redisConfig"`
}

type CommonConfig struct {
	Host               string `yaml:"host"`
	HostName           string `yaml:"hostname"`
	UsedDB             string `yaml:"usedDB"`
	ContentType        string `yaml:"contentType"`
	InnerServiceFormat string `yaml:"innerServiceFormat"`
}

type SecurityConfig struct {
	TLSCertificate     *meta.TLSCertificate `yaml:"tlsCertificate"`
}

type DataStoreConfig struct {
	URL              string `yaml:"URL"`
	TLS              *meta.TLSCertificate `yaml:"tlsCertificate"`
	Database   string `yaml:"database"`
	Collection string `yaml:"collection"`
	Login      string `yaml:"login"`
	Password   string `yaml:"password"`
	Timeout    int    `yaml:"timeout"`
}

type AuthConfig struct {
	PublicKeyPath        string `yaml:"publicKeyPath"`
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
