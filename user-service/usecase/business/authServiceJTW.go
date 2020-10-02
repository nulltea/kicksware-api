package business

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/xid"

	"go.kicksware.com/api/user-service/core/meta"
	"go.kicksware.com/api/user-service/core/model"
	"go.kicksware.com/api/user-service/core/service"
	"go.kicksware.com/api/user-service/env"
)

type authService struct {
	userService service.UserService
	issuerName string
	expirationDelta int
	privateKey *rsa.PrivateKey
	publicKey *rsa.PublicKey
}

func NewAuthServiceJWT(userService service.UserService, authConfig env.AuthConfig) service.AuthService {
	return &authService{
		userService,
		authConfig.IssuerName,
		authConfig.TokenExpirationDelta,
		getPrivateKey(authConfig.PrivateKeyPath),
		getPublicKey(authConfig.PublicKeyPath),
	}
}

func (s *authService) SingUp(user *model.User) (*meta.AuthToken, error) {
	if err := s.userService.Register(user); err != nil {
		return nil, err
	}
	return s.GenerateToken(user)
}

func (s *authService) Login(user *model.User) (*meta.AuthToken, error) {
	registered, err := s.userService.FetchByEmail(user.Email); if err != nil || registered == nil {
		return nil, err
	}

	if registered.PasswordHash != user.PasswordHash {
		return nil, service.ErrPasswordInvalid
	}
	// if !registered.Confirmed {
	// 	return nil, service.ErrNotConfirmed
	// }

	return s.GenerateToken(registered)
}

func (s *authService) Remote(user *model.User) (*meta.AuthToken, error) {
	if user == nil || len(user.UniqueID) == 0 {
		return nil, service.ErrInvalidRemoteID
	} else if len(user.Provider) == 0 || user.Provider == model.Internal {
		return nil, service.ErrInvalidRemoteProvider
	}
	remoteID := user.UniqueID
	provider := user.Provider

	if connected, err := s.userService.FetchRemote(remoteID, provider); err == nil && connected != nil {
		return s.GenerateToken(connected)
	}

	if len(user.Email) != 0 {
		if connected, err := s.userService.FetchByEmail(user.Email); err == nil && connected != nil {
			if err = s.userService.ConnectProvider(connected.UniqueID, remoteID, provider); err != nil {
				return nil, err
			}
			return s.GenerateToken(connected)
		}
	}

	return s.SingUp(user)
}

func (s *authService) Guest() (*meta.AuthToken, error) {
	return s.GenerateToken(&model.User{
		Role: model.Guest,
		UniqueID: xid.New().String(),
	})
}

func (s *authService) GenerateToken(user *model.User) (*meta.AuthToken, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	expiresAt := time.Now().Add(time.Hour * time.Duration(s.expirationDelta))
	token.Claims = &meta.AuthClaims {
		Issuer: s.issuerName,
		UniqueID: user.UniqueID,
		Email: user.Email,
		Role: string(user.Role),
		ExpiresAt: expiresAt.Unix(),
		IssuedAt: time.Now().Unix(),
	}
	tokenString, err := token.SignedString(s.privateKey)
	if err != nil {
		return nil, err
	}
	return meta.NewAuthToken(tokenString, expiresAt), nil
}

func (s *authService) Refresh(raw string) (*meta.AuthToken, error) {
	token, _ := s.parse(raw); if token == nil {
		return s.Guest()
	}
	claims, err := GetClaims(token); if err != nil {
		return s.Guest()
	}

	user, err := s.userService.FetchOne(claims.UniqueID); if err != nil {
		return s.Guest()
	}

	return s.GenerateToken(user)
}


func (s *authService) PublicKey() *rsa.PublicKey {
	return s.publicKey
}

func (s *authService) Logout(token string) error {
	return nil
}

func (s *authService) parse(raw string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(raw, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
			return s.PublicKey(), nil
		}
		return nil, fmt.Errorf("authenticator: unexpected signing method: %q", token.Header["alg"])
	})
	return
}

func getPrivateKey(keyPath string) *rsa.PrivateKey {
	privateKeyFile, err := os.Open(keyPath)
	if err != nil {
		panic(err)
	}

	pemFileInfo, _ := privateKeyFile.Stat()
	var size int64 = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pemBytes)

	data, _ := pem.Decode(pemBytes)

	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS8PrivateKey(data.Bytes); if err != nil {
		panic(err)
	}
	privateKey, ok := privateKeyImported.(*rsa.PrivateKey); if !ok {
		return nil
	}
	return privateKey
}

func getPublicKey(keyPath string) *rsa.PublicKey {
	publicKeyFile, err := os.Open(keyPath)
	if err != nil {
		panic(err)
	}

	pemFileInfo, _ := publicKeyFile.Stat()
	var size int64 = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pemBytes)

	data, _ := pem.Decode(pemBytes)

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes); if err != nil {
		panic(err)
	}

	publicKey, ok := publicKeyImported.(*rsa.PublicKey); if !ok {
		return nil
	}

	return publicKey
}

func GetClaims(token *jwt.Token) (*meta.AuthClaims, error) {
	payload, err := json.Marshal(token.Claims); if err != nil {
		return nil, err
	}
	claims := &meta.AuthClaims{}

	if err = json.Unmarshal(payload, claims); err != nil {
		return nil, err
	}
	return claims, nil
}