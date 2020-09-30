package rest

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/timoth-y/kicksware-api/service-common/util"

	"github.com/timoth-y/kicksware-api/service-common/core/meta"

	"github.com/timoth-y/kicksware-api/order-service/core/model"
	"github.com/timoth-y/kicksware-api/order-service/core/service"
	"github.com/timoth-y/kicksware-api/order-service/env"
	"github.com/timoth-y/kicksware-api/order-service/usecase/business"
	"github.com/timoth-y/kicksware-api/order-service/usecase/serializer/json"
	"github.com/timoth-y/kicksware-api/order-service/usecase/serializer/msg"
)

type Handler struct {
	service     service.OrderService
	auth        service.AuthService
	contentType string
}

func NewHandler(service service.OrderService, auth service.AuthService, config env.CommonConfig) *Handler {
	return &Handler{
		service,
		auth,
		config.ContentType,
	}
}

func (h *Handler) GetOne(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) PostOne(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	order, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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


func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) Count(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) getRequestBody(r *http.Request) (*model.Order, error) {
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


func (h *Handler) getRequestBodies(r *http.Request) ([]*model.Order, error) {
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

func (h *Handler) serializer(contentType string) service.OrderSerializer {
	if contentType == "application/x-msgpack" {
		return msg.NewSerializer()
	}
	return json.NewSerializer()
}