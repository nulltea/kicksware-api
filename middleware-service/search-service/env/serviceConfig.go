package env

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ServiceConfig struct {
	Common  CommonConfig  `yaml:"commonConfig"`
	Auth    AuthConfig    `yaml:"authConfig"`
	Elastic ElasticConfig `yaml:"elasticConfig"`
}

type CommonConfig struct {
	Host               string `yaml:"host"`
	HostName           string `yaml:"hostname"`
	ContentType        string `yaml:"contentType"`
	InnerServiceFormat string `yaml:"innerServiceFormat"`
}

type ElasticConfig struct {
	URL          string `yaml:"URL"`
	Index        string `yaml:"index"`
	StartupDelay int    `yaml:"startupDelay"`
	Sniffing     bool   `yaml:"sniffing"`
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
