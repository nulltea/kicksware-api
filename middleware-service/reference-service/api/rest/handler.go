package rest

import (
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/fatih/structs"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"

	"reference-service/api/common"
	"reference-service/core/model"
	"reference-service/core/service"
	"reference-service/middleware/business"
	"reference-service/middleware/serializer/json"
	"reference-service/middleware/serializer/msg"
)

type RestfulHandler interface {
	GetOne(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	PostQuery(http.ResponseWriter, *http.Request)
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
	params := &common.RequestParams{}
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	params.AssignParams(r)
	h.setupResponse(w, params.ApplyParams(sneakerReferences), http.StatusOK)
}

func (h *handler) PostQuery(w http.ResponseWriter, r *http.Request) {
	query, err := h.getRequestQueryBody(r)
	params := &common.RequestParams{}
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	sneakerReferences, err := h.Service.FetchQuery(query)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	params.AssignParams(r)
	h.setupResponse(w, params.ApplyParams(sneakerReferences), http.StatusOK)
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

func (h *handler) getRequestQueryBody(r *http.Request) (map[string]interface{}, error) {
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

func (h *handler) getRequestParams(r *http.Request) (requestParams common.RequestParams) {
	query := r.URL.Query();
	properties := structs.Names(requestParams)
	prmType := reflect.TypeOf(requestParams);
	for _, prop := range properties {
		value := query.Get(prop)
		paramField := reflect.ValueOf(&requestParams).Elem().FieldByName(prop);
		field, _ := prmType.FieldByName(prop)
		switch field.Type.Kind().String() {
		case "string":
			paramField.SetString(value);
		case "int":
		case "float":
			if num, err := strconv.ParseInt(value, 10, 32); err != nil {
				paramField.SetInt(num);
			}
		case "bool":
			if sign, err := strconv.ParseBool(value); err != nil {
				paramField.SetBool(sign);
			}
		default:
			paramField.SetString(value);
		}
	}
	return
}

func (h *handler) serializer(contentType string) service.SneakerReferenceSerializer {
	if contentType == "application/x-msgpack" {
		return msg.NewSerializer()
	}
	return json.NewSerializer()
}