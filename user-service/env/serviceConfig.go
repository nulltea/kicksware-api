package env

import (
	"io/ioutil"

	"github.com/golang/glog"
	"go.kicksware.com/api/service-common/config"
	"go.kicksware.com/api/service-common/core/meta"
	"gopkg.in/yaml.v2"
)

type ServiceConfig struct {
	Common       config.CommonConfig   `yaml:"commonConfig"`
	Security     config.SecurityConfig `yaml:"securityConfig"`
	Auth         AuthConfig            `yaml:"authConfig"`
	Mail         MailConfig            `yaml:"mailConfig"`
	FallbackMail MailConfig            `yaml:"fallbackMailConfig"`
	Mongo        DataStoreConfig       `yaml:"mongoConfig"`
	Postgres     DataStoreConfig       `yaml:"postgresConfig"`
	Redis        DataStoreConfig       `yaml:"redisConfig"`
	Personal     PersonalConfig        `yaml:"personalConfig"`
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

type PersonalConfig struct {
	SunnyUserEmail    string `yaml:"sunnyUserEmail"`
	SunnyUserIdPrefix string `yaml:"sunnyUserIdPrefix"`
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
