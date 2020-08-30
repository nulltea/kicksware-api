package rest

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/timoth-y/kicksware-platform/middleware-service/service-common/util"

	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/core/meta"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/core/model"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/core/service"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/env"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/usecase/business"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/usecase/serializer/json"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/usecase/serializer/msg"
)

type RestfulHandler interface {
	// Endpoint handlers:
	GetOne(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	PostOne(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Patch(http.ResponseWriter, *http.Request)
	Count(http.ResponseWriter, *http.Request)
	// Middleware:
	Authenticator(next http.Handler) http.Handler
	Authorizer(next http.Handler) http.Handler
	UserSetter(next http.Handler) http.Handler
}

type handler struct {
	service     service.OrderService
	auth        service.AuthService
	contentType string
}

func NewHandler(service service.OrderService, auth service.AuthService, config env.CommonConfig) RestfulHandler {
	return &handler{
		service,
		auth,
		config.ContentType,
	}
}

func (h *handler) GetOne(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r,"orderID")
	params := NewRequestParams(r)
	order, err := h.service.FetchOne(code, params)
	if err != nil {
		if errors.Cause(err) == business.ErrOrderNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, order, http.StatusOK)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	var orders []*model.Order
	var err error
	params := NewRequestParams(r)

	if r.Method == http.MethodPost {
		query, err := h.getRequestQuery(r); if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		orders, err = h.service.FetchQuery(query, params)
	} else if r.Method == http.MethodGet {
		codes := r.URL.Query()["orderID"]
		if codes != nil && len(codes) > 0 {
			orders, err = h.service.Fetch(codes, params)
		} else {
			orders, err = h.service.FetchAll(params)
		}
	}

	if err != nil {
		if errors.Cause(err) == business.ErrOrderNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, orders, http.StatusOK)
}

func (h *handler) PostOne(w http.ResponseWriter, r *http.Request) {
	order, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.StoreOne(order)
	if err != nil {
		if errors.Cause(err) == business.ErrOrderNotValid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, order, http.StatusOK)
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	order, err := h.getRequestBodies(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.service.Store(order)
	if err != nil {
		if errors.Cause(err) == business.ErrOrderNotValid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, order, http.StatusOK)
}


func (h *handler) Patch(w http.ResponseWriter, r *http.Request) {
	order, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.service.Modify(order)
	if err != nil {
		if errors.Cause(err) == business.ErrOrderNotValid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *handler) Count(w http.ResponseWriter, r *http.Request) {
	var count int
	var err error
	params := NewRequestParams(r)

	if r.Method == http.MethodPost {
		query, err := h.getRequestQuery(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		count, err = h.service.Count(query, params)
	} else if r.Method == http.MethodGet {
		query := r.URL.Query()
		if query != nil && len(query) > 0 {
			count, err = h.service.Count(util.ToQueryMap(query), params)
		} else {
			count, err = h.service.CountAll()
		}
	}

	if err != nil {
		if errors.Cause(err) == business.ErrOrderNotFound {
			h.setupResponse(w, 0, http.StatusOK)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, count, http.StatusOK)
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

func (h *handler) getRequestBody(r *http.Request) (*model.Order, error) {
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


func (h *handler) getRequestBodies(r *http.Request) ([]*model.Order, error) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	bodies, err := h.serializer(contentType).DecodeRange(requestBody)
	if err != nil {
		return nil, err
	}
	return bodies, nil
}

func (h *handler) serializer(contentType string) service.OrderSerializer {
	if contentType == "application/x-msgpack" {
		return msg.NewSerializer()
	}
	return json.NewSerializer()
}