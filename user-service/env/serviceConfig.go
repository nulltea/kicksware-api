package env

import (
	"io/ioutil"

	"github.com/golang/glog"
	"go.kicksware.com/api/service-common/config"
	"gopkg.in/yaml.v2"
)

type ServiceConfig struct {
	Common          config.CommonConfig    `yaml:"commonConfig"`
	Security        config.SecurityConfig  `yaml:"securityConfig"`
	UsersDB         config.DataStoreConfig `yaml:"usersDBConfig"`
	LikesDB         config.DataStoreConfig `yaml:"likesDBConfig"`
	RemotesDB       config.DataStoreConfig `yaml:"remotesDBConfig"`
	SubscriptionsDB config.DataStoreConfig `yaml:"subscriptionsDBConfig"`
	Auth            AuthConfig             `yaml:"authConfig"`
	Mail            MailConfig             `yaml:"mailConfig"`
	FallbackMail    MailConfig             `yaml:"fallbackMailConfig"`
	Personal        PersonalConfig         `yaml:"personalConfig"`
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
