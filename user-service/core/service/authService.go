package service

import (
	"crypto/rsa"
	"errors"

	"go.kicksware.com/api/user-service/core/meta"
	"go.kicksware.com/api/user-service/core/model"
)

var (
	ErrPasswordInvalid       = errors.New("user-service/authService: invalid user password")
	ErrNotConfirmed          = errors.New("user-service/authService: user email not confirmed")
	ErrInvalidRemoteID       = errors.New("user-service/authService: invalid remote OAuth identifier")
	ErrInvalidRemoteProvider = errors.New("user-service/authService: invalid remote OAuth provider")
)

type AuthService interface {
	SingUp(user *model.User) (*meta.AuthToken, error)
	Login(user *model.User) (*meta.AuthToken, error)
	Remote(user *model.User) (*meta.AuthToken, error)
	Guest() (*meta.AuthToken, error)
	VerifyAccessKey(hash []byte) bool
	GenerateToken(user *model.User) (*meta.AuthToken, error)
	Refresh(token string) (*meta.AuthToken, error)
	PublicKey() *rsa.PublicKey
	Logout(token string) error
}
