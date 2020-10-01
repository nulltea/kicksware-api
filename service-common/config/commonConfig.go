package config

import "github.com/timoth-y/kicksware-api/service-common/core/meta"

type CommonConfig struct {
	Host               string `yaml:"host"`
	HostName           string `yaml:"hostname"`
	ContentType        string `yaml:"contentType"`
	InnerServiceFormat string `yaml:"innerServiceFormat"`
}

type SecurityConfig struct {
	TLSCertificate     *meta.TLSCertificate `yaml:"tlsCertificate"`
}

type AuthConfig struct {
	PublicKeyPath  string               `yaml:"publicKeyPath"`
	AuthEndpoint   string               `yaml:"authEndpoint"`
	TLSCertificate *meta.TLSCertificate `yaml:"tlsCertificate"`
}