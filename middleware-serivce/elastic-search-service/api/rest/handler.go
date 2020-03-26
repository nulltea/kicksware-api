package rest

import (
	"elastic-search-service/core/model"
	"elastic-search-service/core/service"
	"elastic-search-service/middleware/business"
	"elastic-search-service/middleware/serializer/json"
	"elastic-search-service/middleware/serializer/msg"
	"elastic-search-service/util"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
)

type RestfulHandler interface {
	Get(http.ResponseWriter, *http.Request)
	GetBy(http.ResponseWriter, *http.Request)
	GetSKU(http.ResponseWriter, *http.Request)
	GetBrand(http.ResponseWriter, *http.Request)
	GetModel(http.ResponseWriter, *http.Request)
	PostOne(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	PostAll(http.ResponseWriter, *http.Request)
	PostQuery(http.ResponseWriter, *http.Request)
}

type handler struct {
	Service service.SneakerReferenceService
	ContentType string
}

func NewHandler(service service.SneakerReferenceService, contentType string) RestfulHandler {
	return &handler{
		Service:     service,
		ContentType: contentType,
	}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()["query"][0]
	ref, err := h.Service.Search(query)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, ref, http.StatusOK)
}

func (h *handler) GetBy(w http.ResponseWriter, r *http.Request) {
	field := chi.URLParam(r,"field")
	query := r.URL.Query()["query"][0]
	refs, err := h.Service.SearchBy(field, query)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, refs, http.StatusOK)
}

func (h *handler) GetSKU(w http.ResponseWriter, r *http.Request) {
	sku := chi.URLParam(r, "sku")
	refs, err := h.Service.SearchSKU(sku)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, refs, http.StatusOK)
}

func (h *handler) GetBrand(w http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")
	refs, err := h.Service.SearchBrand(brand)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, refs, http.StatusOK)
}

func (h *handler) GetModel(w http.ResponseWriter, r *http.Request) {
	model := chi.URLParam(r, "model")
	refs, err := h.Service.SearchModel(model)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, refs, http.StatusOK)
}

func (h *handler) PostOne(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	if err := h.Service.SyncOne(code);  err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	codes := r.URL.Query()["codes"]
	if err := h.Service.Sync(codes);  err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *handler) PostAll(w http.ResponseWriter, r *http.Request) {
	if err := h.Service.SyncAll();  err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *handler) PostQuery(w http.ResponseWriter, r *http.Request) {
	query := util.ToQueryMap(r.URL.Query())
	if err := h.Service.SyncQuery(query);  err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
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

func (h *handler) getRequestBody(r *http.Request) (*model.SneakerReference, error) {
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

func (h *handler) serializer(contentType string) service.SneakerReferenceSerializer {
	if contentType == "application/x-msgpack" {
		return msg.NewSerializer()
	}
	return json.NewSerializer()
}