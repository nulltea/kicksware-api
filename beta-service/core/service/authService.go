package service

import (
	"crypto/rsa"
)

type AuthService interface {
	PublicKey() *rsa.PublicKey
}
