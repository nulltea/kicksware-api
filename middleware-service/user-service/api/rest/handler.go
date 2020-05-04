package rest

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"

	"user-service/core/meta"
	"user-service/core/model"
	"user-service/core/service"
	"user-service/middleware/business"
	"user-service/middleware/serializer/json"
	"user-service/middleware/serializer/msg"
)

type RestfulHandler interface {
	GetOne(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Patch(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	SingUp(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	RefreshToken(http.ResponseWriter, *http.Request)
	Logout(http.ResponseWriter, *http.Request)
}

type handler struct {
	Service service.UserService
	Auth service.AuthService
	ContentType string
}

func NewHandler(service service.UserService, auth service.AuthService, contentType string) RestfulHandler {
	return &handler{
		Service:     service,
		Auth:        auth,
		ContentType: contentType,
	}
}

func (h *handler) GetOne(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r,"username")
	user, err := h.Service.FetchOne(username)
	if err != nil {
		if errors.Cause(err) == business.ErrUserNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, user, http.StatusOK)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	var users []*model.User
	var err error

	if r.Method == http.MethodPost {
		query, err := h.getRequestQuery(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		users, err = h.Service.FetchQuery(query)
	} else if r.Method == http.MethodGet {
		if codes := r.URL.Query()["userId"]; codes != nil && len(codes) != 0 {
			users, err = h.Service.Fetch(codes)
		} else {
			users, err = h.Service.FetchAll()
		}
	}

	if err != nil {
		if errors.Cause(err) == business.ErrUserNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, users, http.StatusOK)
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	user, err := h.getRequestBody(r); if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.Service.Register(user)
	if err != nil {
		if errors.Cause(err) == business.ErrUserInvalid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, user, http.StatusOK)
}

func (h *handler) Patch(w http.ResponseWriter, r *http.Request) {
	user, err := h.getRequestBody(r); if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.Service.Modify(user)
	if err != nil {
		if errors.Cause(err) == business.ErrUserInvalid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *handler) Put(w http.ResponseWriter, r *http.Request) {
	user, err := h.getRequestBody(r); if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.Service.Replace(user)
	if err != nil {
		if errors.Cause(err) == business.ErrUserInvalid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r,"sneakerId")
	err := h.Service.Remove(code)
	if err != nil {
		if errors.Cause(err) == business.ErrUserNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *handler) SingUp(w http.ResponseWriter, r *http.Request) {
	user, err := h.getRequestBody(r); if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.Auth.SingUp(user); if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupAuthCookie(w, token)
	h.setupResponse(w, token, http.StatusOK)
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	user, err := h.getRequestBody(r); if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.Auth.Login(user); if err != nil {
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

func (h *handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	user, err := h.getRequestBody(r); if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.Auth.GenerateToken(user); if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupAuthCookie(w, token)
	h.setupResponse(w, token, http.StatusOK)
}

func (h *handler) Logout(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r,"token")
	if err := h.Auth.Logout(token); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.recallAuthCookie(w)
	h.setupResponse(w, token, http.StatusOK)
}

func (h *handler) setupResponse(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", h.ContentType)
	w.WriteHeader(statusCode)
	if body != nil {
		raw, err := h.serializer(h.ContentType).Encode(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(raw); err != nil {
			log.Println(err)
		}
	}
}

func (h *handler) setupAuthCookie(w http.ResponseWriter, token *meta.AuthToken) {
	cookie := &http.Cookie{
		Name: "AuthToken",
		Value: token.Token,
		Expires: token.Expires,
	}
	http.SetCookie(w, cookie)
}

func (h *handler) recallAuthCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name: "AuthToken",
		Expires: time.Now(),
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func (h *handler) getRequestBody(r *http.Request) (*model.User, error) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	body, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (h *handler) getRequestQuery(r *http.Request) (meta.RequestQuery, error) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	body, err := h.serializer(contentType).DecodeMap(requestBody)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (h *handler) serializer(contentType string) service.UserSerializer {
	if contentType == "application/x-msgpack" {
		return msg.NewSerializer()
	}
	return json.NewSerializer()
}