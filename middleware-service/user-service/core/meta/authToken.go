package meta

import "time"

type AuthToken struct {
	Token string
	Success bool
	Expires time.Time
}

func NewAuthToken(token string, expires time.Time) *AuthToken {
	return &AuthToken{
		token,
		true,
		expires,
	}
}