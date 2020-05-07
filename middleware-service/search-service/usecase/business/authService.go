package business

import (
	"crypto/rsa"

	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/util"

	"search-service/core/service"
	"search-service/env"
)

type authService struct {
	publicKey *rsa.PublicKey
}

func NewAuthService(authConfig env.AuthConfig) service.AuthService {
	return &authService{
		util.GetPublicKey(authConfig.PublicKeyPath),
	}
}

func (s *authService) PublicKey() *rsa.PublicKey {
	return s.publicKey
}