package rest

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"

	"product-service/core/model"
	"product-service/core/service"
	"product-service/middleware/business"
	"product-service/middleware/serializer/json"
	"product-service/middleware/serializer/msg"
	"product-service/util"
)

type RestfulHandler interface {
	GetOne(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	GetQuery(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	PutImages(http.ResponseWriter, *http.Request)
	Patch(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type handler struct {
	Service service.SneakerProductService
	ContentType string
}

func NewHandler(service service.SneakerProductService, contentType string) RestfulHandler {
	return &handler{
		Service:     service,
		ContentType: contentType,
	}
}

func (h *handler) GetOne(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r,"sneakerId")
	sneakerProduct, err := h.Service.FetchOne(code)
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

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	codes := r.URL.Query()["sneakerId"]
	sneakerProducts, err := h.Service.Fetch(codes)
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

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	sneakerProducts, err := h.Service.FetchAll()
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

func (h *handler) GetQuery(w http.ResponseWriter, r *http.Request) {
	query := util.ToQueryMap(r.URL.Query())
	sneakerProducts, err := h.Service.FetchQuery(query)
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

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	sneakerProduct, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.Service.Store(sneakerProduct)
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

func (h *handler) PutImages(w http.ResponseWriter, r *http.Request) {
	files, err := h.getRequestFiles(r)
	if err != nil || len(files) == 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	code := chi.URLParam(r,"sneakerId")
	sneakerProduct, err := h.Service.FetchOne(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sneakerProduct.Images = files
	if err = h.Service.Modify(sneakerProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.setupResponse(w, nil, http.StatusOK)
}

func (h *handler) Patch(w http.ResponseWriter, r *http.Request) {
	sneakerProduct, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.Service.Modify(sneakerProduct)
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

func (h *handler) Put(w http.ResponseWriter, r *http.Request) {
	sneakerProduct, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.Service.Replace(sneakerProduct)
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

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r,"sneakerId")
	err := h.Service.Remove(code)
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

func (h *handler) getRequestBody(r *http.Request) (*model.SneakerProduct, error) {
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

func (h *handler) getRequestFiles(r *http.Request) (files map[string][]byte, err error) {
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

func (h *handler) serializer(contentType string) service.SneakerProductSerializer {
	if contentType == "application/x-msgpack" {
		return msg.NewSerializer()
	}
	return json.NewSerializer()
}