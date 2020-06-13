package rest

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/victorspringer/http-cache"

	"cdn-service/core/meta"
	"cdn-service/core/model"
	"cdn-service/core/service"
	"cdn-service/env"
)

type RestfulHandler interface {
	// Endpoint handlers:
	Get(http.ResponseWriter, *http.Request)
	GetCropped(http.ResponseWriter, *http.Request)
	GetResized(http.ResponseWriter, *http.Request)
	GetThumbnail(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	// Middleware:
	Authenticator(next http.Handler) http.Handler
	Authorizer(next http.Handler) http.Handler
	CacheController(next http.Handler) http.Handler
}

type handler struct {
	service     service.ContentService
	auth        service.AuthService
	cache       *cache.Client
}

func NewHandler(service service.ContentService, auth service.AuthService, config env.CommonConfig) RestfulHandler {
	return &handler{
		service,
		auth,
		newCacheClient(),
	}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	query := meta.ContentQuery {
		Filename:   chi.URLParam(r, "filename"),
		Collection: chi.URLParam(r, "collection"),
	}
	if len(query.Collection) == 0 || len(query.Filename) == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	img, err := h.service.Original(query); if err != nil {
		http.Error(w, errors.Wrap(err, http.StatusText(http.StatusInternalServerError)).Error(),
			http.StatusInternalServerError)
		glog.Errorln(err)
		return
	}

	h.setupResponse(w, img, http.StatusOK)
}



func (h *handler) GetCropped(w http.ResponseWriter, r *http.Request) {
	query := meta.ContentQuery {
		Filename:   chi.URLParam(r, "filename"),
		Collection: chi.URLParam(r, "collection"),
	}
	if len(query.Collection) == 0 || len(query.Filename) == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	options := ParseOptions(r); if options.Height == 0 && options.Width == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	img, err := h.service.Crop(query, options); if err != nil {
		http.Error(w, errors.Wrap(err, http.StatusText(http.StatusInternalServerError)).Error(),
			http.StatusInternalServerError)
		glog.Errorln(err)
		return
	}

	h.setupResponse(w, img, http.StatusOK)
}

func (h *handler) GetResized(w http.ResponseWriter, r *http.Request) {
	query := meta.ContentQuery {
		Filename:   chi.URLParam(r, "filename"),
		Collection: chi.URLParam(r, "collection"),
	}
	if len(query.Collection) == 0 || len(query.Filename) == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	options := ParseOptions(r); if options.Height == 0 && options.Width == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	img, err := h.service.Resize(query, options); if err != nil {
		http.Error(w, errors.Wrap(err, http.StatusText(http.StatusInternalServerError)).Error(),
			http.StatusInternalServerError)
		glog.Errorln(err)
		return
	}
	h.setupResponse(w, img, http.StatusOK)
}

func (h *handler) GetThumbnail(w http.ResponseWriter, r *http.Request) {
	query := meta.ContentQuery {
		Filename:   chi.URLParam(r, "filename"),
		Collection: chi.URLParam(r, "collection"),
	}
	if len(query.Collection) == 0 || len(query.Filename) == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	img, err := h.service.Thumbnail(query); if err != nil {
		http.Error(w, errors.Wrap(err, http.StatusText(http.StatusInternalServerError)).Error(),
			http.StatusInternalServerError)
		glog.Errorln(err)
		return
	}
	h.setupResponse(w, img, http.StatusOK)
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (h *handler) setupResponse(w http.ResponseWriter, content *model.Content, statusCode int) {
	w.Header().Set("Content-Length", strconv.Itoa(len(content.Data)))
	w.Header().Set("Content-Type", string(content.MimeType))
	w.WriteHeader(statusCode)
	w.Write(content.Data)
}


