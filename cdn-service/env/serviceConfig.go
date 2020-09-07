package env

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type ServiceConfig struct {
	Common CommonConfig    `yaml:"commonConfig"`
	Auth   AuthConfig      `yaml:"authConfig"`
	Files  DataStoreConfig `yaml:"filesConfig"`
	Mongo  DataStoreConfig `yaml:"mongoConfig"`
	Redis  DataStoreConfig `yaml:"redisConfig"`
}

type CommonConfig struct {
	Host        string `yaml:"host"`
	HostName    string `yaml:"hostname"`
	UsedDB      string `yaml:"usedDB"`
	ContentType string `yaml:"contentType"`
	MaxSize     int
	TailOnly    bool
	ShowInfo    bool
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
	PublicKeyPath string `yaml:"publicKeyPath"`
}

func ReadServiceConfig(filename string) (sc ServiceConfig, err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = yaml.Unmarshal(file, &sc)
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}
