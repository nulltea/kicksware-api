package rest

import (
	"reference-service/core/model"
	"reference-service/core/service"
	"reference-service/middleware/business"
	"reference-service/middleware/serializer/json"
	"reference-service/middleware/serializer/msg"
	"reference-service/util"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
)

type RestfulHandler interface {
	GetOne(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	GetQuery(http.ResponseWriter, *http.Request)
	PostOne(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Patch(http.ResponseWriter, *http.Request)
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

func (h *handler) GetOne(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r,"referenceId")
	sneakerReference, err := h.Service.FetchOne(code)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, sneakerReference, http.StatusOK)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	codes := r.URL.Query()["referenceId"]
	sneakerReferences, err := h.Service.Fetch(codes)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, sneakerReferences, http.StatusOK)
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	sneakerReferences, err := h.Service.FetchAll()
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, sneakerReferences, http.StatusOK)
}

func (h *handler) GetQuery(w http.ResponseWriter, r *http.Request) {
	query := util.ToQueryMap(r.URL.Query())
	sneakerReferences, err := h.Service.FetchQuery(query)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, sneakerReferences, http.StatusOK)
}

func (h *handler) PostOne(w http.ResponseWriter, r *http.Request) {
	sneakerReference, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.Service.StoreOne(sneakerReference)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotValid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, sneakerReference, http.StatusOK)
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	sneakerReference, err := h.getRequestBodies(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.Service.Store(sneakerReference)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotValid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, sneakerReference, http.StatusOK)
}


func (h *handler) Patch(w http.ResponseWriter, r *http.Request) {
	sneakerReference, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.Service.Modify(sneakerReference)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotValid {
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
	body, err := h.serializer(contentType).DecodeOne(requestBody)
	if err != nil {
		return nil, err
	}
	return body, nil
}


func (h *handler) getRequestBodies(r *http.Request) ([]*model.SneakerReference, error) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	bodies, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		return nil, err
	}
	return bodies, nil
}

func (h *handler) serializer(contentType string) service.SneakerReferenceSerializer {
	if contentType == "application/x-msgpack" {
		return msg.NewSerializer()
	}
	return json.NewSerializer()
}