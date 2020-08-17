package meta

type TLSCertificate struct {
	EnableTLS bool `yaml:"enableTLS"`
	CertFile   string `yaml:"certFile"`
	KeyFile    string `yaml:"keyFile"`
}