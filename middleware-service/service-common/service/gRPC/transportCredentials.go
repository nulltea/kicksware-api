package gRPC

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"google.golang.org/grpc/credentials"

	"github.com/timoth-y/kicksware-platform/middleware-service/service-common/core/meta"
)

func LoadServerTLSCredentials(cert *meta.TLSCertificate) (credentials.TransportCredentials, error) {
	pemClientCA, err := ioutil.ReadFile(cert.CACertFile); if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool(); if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	serverCert, err := tls.LoadX509KeyPair(cert.CertFile, cert.KeyFile); if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}
	return credentials.NewTLS(config), nil
}

func LoadClientTLSCredentials(cert *meta.TLSCertificate) (credentials.TransportCredentials, error) {
	pemServerCA, err := ioutil.ReadFile(cert.CACertFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	config := &tls.Config{
		RootCAs:      certPool,
	}
	return credentials.NewTLS(config), nil
}