package meta

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthClaims struct {
	ExpiresAt int64  `json:"exp,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	Subject   string `json:"sub,omitempty"`
	UniqueID  string `json:"nameid,omitempty"`
	Username  string `json:"unique_name,omitempty"`
	Email     string `json:"sub,omitempty"`
	Role      string `json:"role,omitempty"`
}

func (c AuthClaims) Valid() error {
	sc := jwt.StandardClaims{}
	vErr := new(jwt.ValidationError)
	now := jwt.TimeFunc().Unix()

	if sc.VerifyExpiresAt(now, false) == false {
		delta := time.Unix(now, 0).Sub(time.Unix(c.ExpiresAt, 0))
		vErr.Inner = fmt.Errorf("token is expired by %v", delta)
		vErr.Errors |= jwt.ValidationErrorExpired
	}

	if sc.VerifyIssuedAt(now, false) == false {
		vErr.Inner = fmt.Errorf("token used before issued")
		vErr.Errors |= jwt.ValidationErrorIssuedAt
	}

	if sc.VerifyNotBefore(now, false) == false {
		vErr.Inner = fmt.Errorf("token is not valid yet")
		vErr.Errors |= jwt.ValidationErrorNotValidYet
	}

	if vErr == nil {
		return nil
	}
	return vErr
}

