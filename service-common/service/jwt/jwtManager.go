package jwt

import (
	"crypto/rsa"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/timoth-y/kicksware-api/user-service/core/meta"
)

type TokenManager struct {
	PublicKey     *rsa.PublicKey
}

func NewJWTManager(pb *rsa.PublicKey) *TokenManager {
	return &TokenManager{
		PublicKey: pb,
	}
}

func (m *TokenManager) Verify(accessToken string) (*meta.AuthClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&meta.AuthClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
				return m.PublicKey, nil
			}
			return nil, fmt.Errorf("JWTManager: unexpected signing method: %q", token.Header["alg"])
		},
	); if err != nil {
		return nil, fmt.Errorf("access token is invalid: %w", err)
	}

	if claims, ok := token.Claims.(*meta.AuthClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}