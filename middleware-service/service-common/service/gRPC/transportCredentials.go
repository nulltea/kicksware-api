package gRPC

import (
	"google.golang.org/grpc/credentials"

	"github.com/timoth-y/kicksware-platform/middleware-service/service-common/core/meta"
)

func LoadServerTLSCredentials(cert *meta.TLSCertificate) (credentials.TransportCredentials, error) {
	return credentials.NewServerTLSFromFile(cert.CertFile, cert.KeyFile)
}

func LoadClientTLSCredentials(cert *meta.TLSCertificate) (credentials.TransportCredentials, error) {
	return credentials.NewClientTLSFromFile(cert.CACertFile, "")
}