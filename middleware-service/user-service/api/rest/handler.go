package rest

import (
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"user-service/core/model"
	"user-service/core/service"
	"user-service/middleware/business"
	"user-service/middleware/serializer/json"
	"user-service/middleware/serializer/msg"
	"user-service/util"
)

type RestfulHandler interface {
	GetOne(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	GetQuery(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Patch(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type handler struct {
	Service service.UserService
	ContentType string
}

func NewHandler(service service.UserService, contentType string) RestfulHandler {
	return &handler{
		Service:     service,
		ContentType: contentType,
	}
}

func (h *handler) GetOne(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r,"sneakerId")
	user, err := h.Service.FetchOne(code)
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
	codes := r.URL.Query()["sneakerId"]
	users, err := h.Service.Fetch(codes)
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

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.FetchAll()
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

func (h *handler) GetQuery(w http.ResponseWriter, r *http.Request) {
	query := util.ToQueryMap(r.URL.Query())
	users, err := h.Service.FetchQuery(query)
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
	user, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	user, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	user, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

func (h *handler) serializer(contentType string) service.UserSerializer {
	if contentType == "application/x-msgpack" {
		return msg.NewSerializer()
	}
	return json.NewSerializer()
}