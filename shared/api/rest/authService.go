package rest

import (
	"bytes"
	"crypto/md5"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/glog"

	"go.kicksware.com/api/shared/config"
	"go.kicksware.com/api/shared/core"
	"go.kicksware.com/api/shared/util"
)

type authService struct {
	publicKey *rsa.PublicKey
	authEndpoint string
	accessChecksum []byte
}

func NewAuthService(authConfig config.AuthConfig) core.AuthService {
	return &authService{
		util.GetPublicKey(authConfig.PublicKeyPath),
		authConfig.AuthEndpoint,
		getAccessChecksum(authConfig.AccessKey),
	}
}

func (s *authService) PublicKey() *rsa.PublicKey {
	return s.publicKey
}

func (s *authService) Authenticate() (string, error)  {
	resp, err := http.Get(s.guestEndpoint()); if err != nil {
		return "", errors.New("error has occurred in authenticating service: %q\", err")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body); if err != nil {
		return "", err
	}
	token := make(map[string]interface{})
	err = json.Unmarshal(bytes, &token); if err != nil {
		return "", err
	}
	return token["Token"].(string), nil
}

func (s *authService) AccessKey() []byte {
	return s.accessChecksum
}


func (s *authService) VerifyAccessKey(hash []byte) bool {
	ok := bytes.Equal(s.accessChecksum, hash)
	return ok
}

func getAccessChecksum(accessKey string) []byte {
	input := strings.NewReader(accessKey)
	hash := md5.New()
	if _, err := io.Copy(hash, input); err != nil {
		glog.Fatalln(err)
	}
	return hash.Sum(nil)
}

func (s *authService) guestEndpoint() string {
	req, err := http.NewRequest("GET", s.authEndpoint, nil); if err != nil {
		glog.Fatalf("error has occurred in core.AuthService.guestEndpoint(): %q", err)
		return ""
	}
	query := req.URL.Query()
	query.Add("access", string(s.accessChecksum))
	req.URL.RawQuery = query.Encode()
	return req.URL.String()
}
