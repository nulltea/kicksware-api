package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"

	"go.kicksware.com/api/user-service/core/meta"
	"go.kicksware.com/api/user-service/core/model"
	"go.kicksware.com/api/user-service/core/service"
	"go.kicksware.com/api/user-service/usecase/business"
)

var (
	ErrInvalidTokenClaims = errors.New("invalid token claims")
)

func (h *Handler) SingUp(w http.ResponseWriter, r *http.Request) {
	user, err := h.getRequestBody(r); if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.auth.SingUp(user); if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupAuthCookie(w, token)
	h.setupResponse(w, token, http.StatusOK)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	user, err := h.getRequestBody(r); if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.auth.Login(user); if err != nil {
		if errors.Cause(err) == service.ErrPasswordInvalid ||
			errors.Cause(err) == service.ErrNotConfirmed {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupAuthCookie(w, token)
	h.setupResponse(w, token, http.StatusOK)
}

func (h *Handler) Remote(w http.ResponseWriter, r *http.Request) {
	user, err := h.getRequestBody(r); if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.auth.Remote(user); if err != nil {
		if errors.Cause(err) == service.ErrInvalidRemoteID {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		} else if errors.Cause(err) == service.ErrInvalidRemoteProvider {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupAuthCookie(w, token)
	h.setupResponse(w, token, http.StatusOK)
}

func (h *Handler) Guest(w http.ResponseWriter, r *http.Request) {
	token, err := h.auth.Guest(); if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupAuthCookie(w, token)
	h.setupResponse(w, token, http.StatusOK)
}

func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	token, err := h.auth.Refresh(chi.URLParam(r,"token")); if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupAuthCookie(w, token)
	h.setupResponse(w, token, http.StatusOK)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r,"token")
	if err := h.auth.Logout(token); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.recallAuthCookie(w)
	h.setupResponse(w, token, http.StatusOK)
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

		claims, err := business.GetClaims(token); if err != nil  {
			http.Error(w, ErrInvalidTokenClaims.Error(), http.StatusInternalServerError)
			fmt.Println()
			return
		}
		if claims != nil && claims.Role != string(model.Guest) {
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

func getClaims(token *jwt.Token) (*meta.AuthClaims, error) {
	payload, err := json.Marshal(token.Claims); if err != nil {
		return nil, err
	}
	claims := &meta.AuthClaims{}

	if err = json.Unmarshal(payload, claims); err != nil {
		return nil, err
	}
	return claims, nil
}
