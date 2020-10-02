package business

import (
	"crypto/rsa"

	"github.com/timoth-y/kicksware-api/service-common/util"

	"go.kicksware.com/api/beta/core/service"
	"go.kicksware.com/api/beta/env"
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