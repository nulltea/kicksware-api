package env

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ServiceConfig struct {
	Common  CommonConfig  `yaml:"commonConfig"`
	Auth    AuthConfig    `yaml:"authConfig"`
	Elastic ElasticConfig `yaml:"elasticConfig"`
	Search  SearchConfig  `yaml:"searchConfig"`
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
	PublicKeyPath string `yaml:"publicKeyPath"`
	AuthEndpoint string `yaml:"authEndpoint"`
}

type SearchConfig struct {
	Type      string   `yaml:"type"`
	Fuzziness string   `yaml:"fuzziness"`
	Slop      int      `yaml:"slop"`
	Fields    []string `yaml:"Fields"`
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
