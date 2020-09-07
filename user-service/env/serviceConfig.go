package env

import (
	"io/ioutil"

	"github.com/golang/glog"
	"github.com/timoth-y/kicksware-platform/middleware-service/service-common/core/meta"
	"gopkg.in/yaml.v2"
)

type ServiceConfig struct {
	Common       CommonConfig    `yaml:"commonConfig"`
	Security     SecurityConfig  `yaml:"securityConfig"`
	Auth         AuthConfig      `yaml:"authConfig"`
	Mail         MailConfig      `yaml:"mailConfig"`
	FallbackMail MailConfig      `yaml:"fallbackMailConfig"`
	Mongo        DataStoreConfig `yaml:"mongoConfig"`
	Postgres     DataStoreConfig `yaml:"postgresConfig"`
	Redis        DataStoreConfig `yaml:"redisConfig"`
}

type CommonConfig struct {
	Host               string               `yaml:"host"`
	HostName           string               `yaml:"hostname"`
	UsedDB             string               `yaml:"usedDB"`
	ContentType        string               `yaml:"contentType"`
	InnerServiceFormat string               `yaml:"innerServiceFormat"`
}

type SecurityConfig struct {
	TLSCertificate     *meta.TLSCertificate `yaml:"tlsCertificate"`
}

type DataStoreConfig struct {
	URL              string `yaml:"URL"`
	TLS              *meta.TLSCertificate `yaml:"TLS"`
	Database         string `yaml:"database"`
	Collection       string `yaml:"collection"`
	LikesCollection  string `yaml:"likesCollection"`
	RemoteCollection string `yaml:"remoteCollection"`
	Login            string `yaml:"login"`
	Password         string `yaml:"password"`
	Timeout          int    `yaml:"timeout"`
}

type AuthConfig struct {
	IssuerName           string `yaml:"issuerName"`
	TokenExpirationDelta int    `yaml:"tokenExpirationDelta"`
	PrivateKeyPath       string `yaml:"privateKeyPath"`
	PublicKeyPath        string `yaml:"publicKeyPath"`
}

type MailConfig struct {
	Server                string `yaml:"server"`
	Address               string `yaml:"address"`
	Password              string `yaml:"password"`
	VerifyEmailTemplate   string `yaml:"verifyEmailTemplate"`
	ResetPasswordTemplate string `yaml:"resetPasswordTemplate"`
	NotificationTemplate  string `yaml:"notificationTemplate"`
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
