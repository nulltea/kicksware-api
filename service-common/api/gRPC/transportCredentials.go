package gRPC

import (
	"google.golang.org/grpc/credentials"

	"go.kicksware.com/api/service-common/core/meta"
)

func LoadServerTLSCredentials(cert *meta.TLSCertificate) (credentials.TransportCredentials, error) {
	return credentials.NewServerTLSFromFile(cert.CertFile, cert.KeyFile)
}

func LoadClientTLSCredentials(cert *meta.TLSCertificate) (credentials.TransportCredentials, error) {
	return credentials.NewClientTLSFromFile(cert.CertFile, "")
}