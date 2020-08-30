package rest

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"

	"github.com/timoth-y/kicksware-platform/middleware-service/reference-service/core/model"

	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/meta"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/service"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/env"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/usecase/business"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/usecase/serializer/json"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/usecase/serializer/msg"
)

type Handler struct {
	search      service.ReferenceSearchService
	sync        service.ReferenceSyncService
	auth        service.AuthService
	contentType string
}

func NewHandler(search service.ReferenceSearchService, sync service.ReferenceSyncService, auth service.AuthService, config env.CommonConfig) *Handler {
	return &Handler{
		search,
		sync,
		auth,
		config.ContentType,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()["query"][0]
	params := NewRequestParams(r)

	ref, err := h.search.Search(query, params)
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

func (h *Handler) GetBy(w http.ResponseWriter, r *http.Request) {
	field := chi.URLParam(r,"field")
	query := r.URL.Query()["query"][0]
	params := NewRequestParams(r)

	refs, err := h.search.SearchBy(field, query, params)
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

func (h *Handler) GetSKU(w http.ResponseWriter, r *http.Request) {
	sku := chi.URLParam(r, "sku")
	params := NewRequestParams(r)

	refs, err := h.search.SearchSKU(sku, params)
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

func (h *Handler) GetBrand(w http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")
	params := NewRequestParams(r)

	refs, err := h.search.SearchBrand(brand, params)
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

func (h *Handler) GetModel(w http.ResponseWriter, r *http.Request) {
	model := chi.URLParam(r, "model")
	params := NewRequestParams(r)

	refs, err := h.search.SearchModel(model, params)
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

func (h *Handler) PostOne(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "referenceId")
	if err := h.sync.SyncOne(code);  err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	codes := r.URL.Query()["referenceId"]
	params := NewRequestParams(r)
	if err := h.sync.Sync(codes, params);  err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *Handler) PostAll(w http.ResponseWriter, r *http.Request) {
	params := NewRequestParams(r)
	if err := h.sync.SyncAll(params);  err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *Handler) PostQuery(w http.ResponseWriter, r *http.Request) {
	query, err := h.getRequestQuery(r); if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params := NewRequestParams(r)
	if err := h.sync.SyncQuery(query, params);  err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
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

func (h *Handler) getRequestBody(r *http.Request) (*model.SneakerReference, error) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	body, err := h.serializer(contentType).DecodeReference(requestBody)
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

func (h *Handler) serializer(contentType string) service.SneakerSearchSerializer {
	if contentType == "application/x-msgpack" {
		return msg.NewSerializer()
	}
	return json.NewSerializer()
}