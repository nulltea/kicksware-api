package rest

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/timoth-y/kicksware-api/service-common/util"

	"github.com/timoth-y/kicksware-api/service-common/core/meta"
	"github.com/timoth-y/kicksware-api/product-service/core/model"
	"github.com/timoth-y/kicksware-api/product-service/core/service"
	"github.com/timoth-y/kicksware-api/product-service/env"
	"github.com/timoth-y/kicksware-api/product-service/usecase/business"
	"github.com/timoth-y/kicksware-api/product-service/usecase/serializer/json"
	"github.com/timoth-y/kicksware-api/product-service/usecase/serializer/msg"
)

type Handler struct {
	service     service.SneakerProductService
	auth        service.AuthService
	contentType string
}

func NewHandler(service service.SneakerProductService, auth service.AuthService, config env.CommonConfig) *Handler {
	return &Handler{
		service,
		auth,
		config.ContentType,
	}
}

func (h *Handler) GetOne(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r,"sneakerId")
	sneakerProduct, err := h.service.FetchOne(code)
	if err != nil {
		if errors.Cause(err) == business.ErrProductNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, sneakerProduct, http.StatusOK)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	var sneakerProducts []*model.SneakerProduct
	var err error
	params := NewRequestParams(r)

	if r.Method == http.MethodPost {
		query, err := h.getRequestQuery(r); if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		sneakerProducts, err = h.service.FetchQuery(query, params)
	} else if r.Method == http.MethodGet {
		codes := r.URL.Query()["sneakerId"]
		if codes != nil && len(codes) > 0 {
			sneakerProducts, err = h.service.Fetch(codes, params)
		} else {
			sneakerProducts, err = h.service.FetchAll(params)
		}
	}

	if err != nil {
		if errors.Cause(err) == business.ErrProductNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, sneakerProducts, http.StatusOK)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	sneakerProduct, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	params := NewRequestParams(r)
	err = h.service.Store(sneakerProduct, params)
	if err != nil {
		if errors.Cause(err) == business.ErrProductInvalid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, sneakerProduct, http.StatusOK)
}

func (h *Handler) PutImages(w http.ResponseWriter, r *http.Request) {
	files, err := h.getRequestFiles(r)
	if err != nil || len(files) == 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	code := chi.URLParam(r,"sneakerId")
	sneakerProduct, err := h.service.FetchOne(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sneakerProduct.Images = files
	if err = h.service.Modify(sneakerProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.setupResponse(w, nil, http.StatusOK)
}

func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
	sneakerProduct, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.service.Modify(sneakerProduct)
	if err != nil {
		if errors.Cause(err) == business.ErrProductInvalid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	sneakerProduct, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.service.Replace(sneakerProduct)
	if err != nil {
		if errors.Cause(err) == business.ErrProductInvalid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r,"sneakerId")
	err := h.service.Remove(code)
	if err != nil {
		if errors.Cause(err) == business.ErrProductNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
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

func (h *Handler) getRequestBody(r *http.Request) (*model.SneakerProduct, error) {
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

func (h *Handler) getRequestFiles(r *http.Request) (files map[string][]byte, err error) {
	files = map[string][]byte{}
	if err := r.ParseMultipartForm(32 << 20 ); err != nil {
		return nil, err
	}
	filesMap := r.MultipartForm.File
	for _, fh := range filesMap {
		for _, h := range fh {
			f, err := h.Open()
			if err != nil {
				continue
			}
			bytes, err := ioutil.ReadAll(f)
			if err != nil {
				continue
			}
			files[h.Filename] = bytes
		}
	}
	return
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
		if errors.Cause(err) == business.ErrProductNotFound {
			h.setupResponse(w, 0, http.StatusOK)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, count, http.StatusOK)
}

func (h *Handler) getRequestQueryBody(r *http.Request) (map[string]interface{}, error) {
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

func (h *Handler) serializer(contentType string) service.SneakerProductSerializer {
	if contentType == "application/x-msgpack" {
		return msg.NewSerializer()
	}
	return json.NewSerializer()
}