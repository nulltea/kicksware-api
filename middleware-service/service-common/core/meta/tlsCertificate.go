package meta

type TLSCertificate struct {
	CACertFile string `yaml:"caCertFile"`
	CertFile   string `yaml:"certFile"`
	KeyFile    string `yaml:"keyFile"`
}