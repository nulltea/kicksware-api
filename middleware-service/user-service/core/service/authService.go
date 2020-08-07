package service

import (
	"crypto/rsa"
	"errors"

	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/meta"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/model"
)

var (
	ErrPasswordInvalid       = errors.New("github.com/timoth-y/kicksware-platform/middleware-service/user-service/authService: invalid user password")
	ErrNotConfirmed          = errors.New("github.com/timoth-y/kicksware-platform/middleware-service/user-service/authService: user email not confirmed")
	ErrInvalidRemoteID       = errors.New("github.com/timoth-y/kicksware-platform/middleware-service/user-service/authService: invalid remote OAuth identifier")
	ErrInvalidRemoteProvider = errors.New("github.com/timoth-y/kicksware-platform/middleware-service/user-service/authService: invalid remote OAuth provider")
)

type AuthService interface {
	SingUp(user *model.User) (*meta.AuthToken, error)
	Login(user *model.User) (*meta.AuthToken, error)
	Remote(user *model.User) (*meta.AuthToken, error)
	Guest() (*meta.AuthToken, error)
	GenerateToken(user *model.User) (*meta.AuthToken, error)
	Refresh(token string) (*meta.AuthToken, error)
	PublicKey() *rsa.PublicKey
	Logout(token string) error
}
