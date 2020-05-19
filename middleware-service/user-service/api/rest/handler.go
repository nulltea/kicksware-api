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
	"user-service/env"
	"user-service/usecase/business"
	"user-service/usecase/serializer/json"
	"user-service/usecase/serializer/msg"
)

type RestfulHandler interface {
	// Endpoint handlers:
	// CRUD
	GetOne(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Patch(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	// Auth
	SingUp(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	Guest(http.ResponseWriter, *http.Request)
	RefreshToken(http.ResponseWriter, *http.Request)
	Logout(http.ResponseWriter, *http.Request)
	// Mail
	SendEmailConfirmation(http.ResponseWriter, *http.Request)
	SendResetPassword(http.ResponseWriter, *http.Request)
	SendNotification(http.ResponseWriter, *http.Request)
	// Interaction
	Like(http.ResponseWriter, *http.Request)
	Unlike(http.ResponseWriter, *http.Request)
	// Middleware:
	Authenticator(next http.Handler) http.Handler
	Authorizer(next http.Handler) http.Handler
}

type handler struct {
	service     service.UserService
	auth        service.AuthService
	mail        service.MailService
	interact    service.InteractService
	contentType string
}

func NewHandler(service service.UserService, auth service.AuthService, mail service.MailService,
	interact service.InteractService, config env.CommonConfig) RestfulHandler {
	return &handler{
		service,
		auth,
		mail,
		interact,
		config.ContentType,
	}
}

func (h *handler) GetOne(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r,"userID")
	user, err := h.service.FetchOne(userID)
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
	params := NewRequestParams(r)

	if r.Method == http.MethodPost {
		query, err := h.getRequestQuery(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		users, err = h.service.FetchQuery(query, params)
	} else if r.Method == http.MethodGet {
		if codes := r.URL.Query()["userID"]; codes != nil && len(codes) != 0 {
			users, err = h.service.Fetch(codes, params)
		} else {
			users, err = h.service.FetchAll(params)
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
	err = h.service.Register(user)
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
	err = h.service.Modify(user)
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
	err = h.service.Replace(user)
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
	code := chi.URLParam(r,"userID")
	err := h.service.Remove(code)
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

func (h *handler) setupResponse(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", h.contentType)
	w.WriteHeader(statusCode)
	if body != nil {
		raw, err := h.serializer(h.contentType).Encode(body)
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