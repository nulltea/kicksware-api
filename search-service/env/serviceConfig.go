package env

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"github.com/golang/glog"
	"github.com/timoth-y/kicksware-api/service-common/core/meta"
)

type ServiceConfig struct {
	Common   CommonConfig   `yaml:"commonConfig"`
	Security SecurityConfig `yaml:"securityConfig"`
	Auth     AuthConfig     `yaml:"authConfig"`
	Elastic  ElasticConfig  `yaml:"elasticConfig"`
	Search   SearchConfig   `yaml:"searchConfig"`
}

type CommonConfig struct {
	Host               string `yaml:"host"`
	HostName           string `yaml:"hostname"`
	ContentType        string `yaml:"contentType"`
	InnerServiceFormat string `yaml:"innerServiceFormat"`
}

type SecurityConfig struct {
	TLSCertificate     *meta.TLSCertificate `yaml:"tlsCertificate"`
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
	TLSCertificate     *meta.TLSCertificate `yaml:"tlsCertificate"`
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
