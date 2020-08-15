package meta

type TLSCertificate struct {
	EnableTLS bool `yaml:"enableTLS"`
	CACertFile string `yaml:"caCertFile"`
	CertFile   string `yaml:"certFile"`
	KeyFile    string `yaml:"keyFile"`
}