package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var (
	errInvalidTokenClaims = errors.New("invalid token claims")
	guestRole = "gst"
)

type authClaims struct {
	UniqueID  string `json:"nameid,omitempty"`
	Username  string `json:"unique_name,omitempty"`
	Email     string `json:"sub,omitempty"`
	Role      string `json:"role,omitempty"`
}

func (h *Handler) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := h.getRequestToken(r); if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			fmt.Println()
			return
		}

		if token == nil || !token.Valid {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) Authorizer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := h.getRequestToken(r); if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			fmt.Println()
			return
		}

		claims, err := getClaims(token); if err != nil {
			http.Error(w, errInvalidTokenClaims.Error(), http.StatusInternalServerError)
			fmt.Println()
			return
		}
		if claims != nil && claims.Role != guestRole {
			r.URL.User = url.User(claims.UniqueID)
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			fmt.Println()
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) getRequestToken(r *http.Request) (token *jwt.Token, err error) {
	token, err = request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
			return h.auth.PublicKey(), nil
		}
		return nil, fmt.Errorf("authenticator: unexpected signing method: %q", token.Header["alg"])
	})
	return
}

func getClaims(token *jwt.Token) (*authClaims, error) {
	payload, err := json.Marshal(token.Claims); if err != nil {
		return nil, err
	}
	claims := &authClaims{}

	if err = json.Unmarshal(payload, claims); err != nil {
		return nil, err
	}
	return claims, nil
}

