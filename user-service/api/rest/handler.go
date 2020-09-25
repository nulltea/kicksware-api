package rest

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"

	"github.com/timoth-y/kicksware-api/user-service/core/meta"
	"github.com/timoth-y/kicksware-api/user-service/core/model"
	"github.com/timoth-y/kicksware-api/user-service/core/service"
	"github.com/timoth-y/kicksware-api/user-service/env"
	"github.com/timoth-y/kicksware-api/user-service/usecase/business"
	"github.com/timoth-y/kicksware-api/user-service/usecase/serializer/json"
	"github.com/timoth-y/kicksware-api/user-service/usecase/serializer/msg"
)

type Handler struct {
	service     service.UserService
	auth        service.AuthService
	mail        service.MailService
	interact    service.InteractService
	contentType string
}

func NewHandler(service service.UserService, auth service.AuthService, mail service.MailService,
	interact service.InteractService, config env.CommonConfig) *Handler {
	return &Handler{
		service,
		auth,
		mail,
		interact,
		config.ContentType,
	}
}

func (h *Handler) GetOne(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) GetTheme(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r,"userID")
	user, err := h.service.FetchOne(code)
	if err != nil {
		if errors.Cause(err) == business.ErrUserNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, user.Settings.Theme, http.StatusOK)
}

func (h *Handler) setupResponse(w http.ResponseWriter, body interface{}, statusCode int) {
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

func (h *Handler) setupAuthCookie(w http.ResponseWriter, token *meta.AuthToken) {
	cookie := &http.Cookie{
		Name: "AuthToken",
		Value: token.Token,
		Expires: token.Expires,
	}
	http.SetCookie(w, cookie)
}

func (h *Handler) recallAuthCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name: "AuthToken",
		Expires: time.Now(),
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func (h *Handler) getRequestBody(r *http.Request) (*model.User, error) {
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

func (h *Handler) getRequestQuery(r *http.Request) (meta.RequestQuery, error) {
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

func (h *Handler) serializer(contentType string) service.UserSerializer {
	if contentType == "application/x-msgpack" {
		return msg.NewSerializer()
	}
	return json.NewSerializer()
}