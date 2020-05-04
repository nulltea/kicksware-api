package business

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"user-service/core/meta"
	"user-service/core/model"
	"user-service/core/service"
)

type authService struct {
	userService service.UserService
	expirationDelta int
}

func NewAuthServiceJWT(userService service.UserService, expirationDelta int) service.AuthService {
	return &authService{
		userService,
		expirationDelta,
	}
}

func (s *authService) SingUp(user *model.User) (*meta.AuthToken, error) {
	if err := s.userService.Register(user); err != nil {
		return nil, err
	}
	return s.GenerateToken(user)
}

func (s *authService) Login(user *model.User)(*meta.AuthToken, error) {
	registered, err := s.userService.FetchOne(user.Username); if err != nil {
		return nil, err
	}

	if registered.PasswordHash != user.PasswordHash {
		return nil, service.ErrPasswordInvalid
	}
	if !registered.Confirmed {
		return nil, service.ErrNotConfirmed
	}

	return s.GenerateToken(user)
}

func (s *authService) GenerateToken(user *model.User) (*meta.AuthToken, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	expiresAt := time.Now().Add(time.Hour * time.Duration(s.expirationDelta))
	token.Claims = &jwt.StandardClaims {
		ExpiresAt: expiresAt.Unix(),
		IssuedAt: time.Now().Unix(),
		Issuer: user.UniqueId,
	}
	tokenString, err := token.SigningString()
	if err != nil {
		return nil, err
	}
	return meta.NewAuthToken(tokenString, expiresAt), nil
}

func (s *authService) Logout(token string) error {
	panic("implement me")
}
