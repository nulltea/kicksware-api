package env

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ServiceConfig struct {
	Common      CommonConfig    `yaml:"commonConfig"`
	Auth        AuthConfig      `yaml:"authConfig"`
	Mongo       DataStoreConfig `yaml:"mongoConfig"`
	Postgres    DataStoreConfig `yaml:"postgresConfig"`
	Redis       DataStoreConfig `yaml:"redisConfig"`
}

type CommonConfig struct {
	Host               string `yaml:"host"`
	HostName           string `yaml:"hostname"`
	UsedDB             string `yaml:"usedDB"`
	ContentType        string `yaml:"contentType"`
	InnerServiceFormat string `yaml:"innerServiceFormat"`
}

type DataStoreConfig struct {
	URL        string `yaml:"URL"`
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
		return
	}
	err = yaml.Unmarshal(file, &sc); if err != nil {
		return
	}
	return
}
