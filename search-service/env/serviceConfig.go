package env

import (
	"io/ioutil"

	"go.kicksware.com/api/service-common/config"
	"gopkg.in/yaml.v2"

	"github.com/golang/glog"
)

type ServiceConfig struct {
	Common   config.CommonConfig   `yaml:"commonConfig"`
	Security config.SecurityConfig `yaml:"securityConfig"`
	Auth     config.AuthConfig     `yaml:"authConfig"`
	Elastic  ElasticConfig         `yaml:"elasticConfig"`
	Search   SearchConfig          `yaml:"searchConfig"`
}

type ElasticConfig struct {
	URL          string `yaml:"URL"`
	Index        string `yaml:"index"`
	StartupDelay int    `yaml:"startupDelay"`
	Sniffing     bool   `yaml:"sniffing"`
}

type SearchConfig struct {
	Type      string   `yaml:"type"`
	Fuzziness string   `yaml:"fuzziness"`
	Slop      int      `yaml:"slop"`
	Fields    []string `yaml:"Fields"`
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
