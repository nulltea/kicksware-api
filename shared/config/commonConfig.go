package config

import "go.kicksware.com/api/shared/core/meta"

type CommonConfig struct {
	Host              string           `yaml:"host"`
	HostName          string           `yaml:"hostname"`
	ContentType       string           `yaml:"contentType"`
	UsedDB            string           `yaml:"usedDB"`
	ApiEndpointFormat string           `yaml:"apiEndpointFormat"`
	RpcEndpointFormat string           `yaml:"rpcEndpointFormat"`
	AmqpConnection    ConnectionConfig `yaml:"amqpConnection"`
}

type ConnectionConfig struct {
	URL string               `yaml:"URL"`
	TLS *meta.TLSCertificate `yaml:"TLS"`
}

type SecurityConfig struct {
	TLSCertificate *meta.TLSCertificate `yaml:"tlsCertificate"`
}

type AuthConfig struct {
	PublicKeyPath  string               `yaml:"publicKeyPath"`
	AuthEndpoint   string               `yaml:"authEndpoint"`
	TLSCertificate *meta.TLSCertificate `yaml:"tlsCertificate"`
	AccessKey      string               `yaml:"accessKey"`
}

type DataStoreConfig struct {
	URL        string               `yaml:"URL"`
	TLS        *meta.TLSCertificate `yaml:"TLS"`
	Database   string               `yaml:"database"`
	Collection string               `yaml:"collection"`
	Login      string               `yaml:"login"`
	Password   string               `yaml:"password"`
	Timeout    int                  `yaml:"timeout"`
}
