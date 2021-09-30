package rest

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"go.kicksware.com/api/shared/api/rest"
	"go.kicksware.com/api/shared/core"

	"github.com/victorspringer/http-cache"

	"go.kicksware.com/api/services/cdn/core/meta"
	"go.kicksware.com/api/services/cdn/core/model"
	"go.kicksware.com/api/services/cdn/core/service"
)

type Handler struct {
	service service.ContentService
	auth    *rest.AuthMiddleware
	cache       *cache.Client
}

func NewHandler(service service.ContentService, auth core.AuthService) *Handler {
	return &Handler{
		service,
		rest.NewAuthMiddleware(auth),
		newCacheClient(),
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	query := meta.ContentQuery{
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



func (h *Handler) GetCropped(w http.ResponseWriter, r *http.Request) {
	query := meta.ContentQuery{
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

func (h *Handler) GetResized(w http.ResponseWriter, r *http.Request) {
	query := meta.ContentQuery{
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

func (h *Handler) GetThumbnail(w http.ResponseWriter, r *http.Request) {
	query := meta.ContentQuery{
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

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (h *Handler) setupResponse(w http.ResponseWriter, content *model.Content, statusCode int) {
	w.Header().Set("Content-Length", strconv.Itoa(len(content.Data)))
	w.Header().Set("Content-Type", string(content.MimeType))
	w.WriteHeader(statusCode)
	w.Write(content.Data)
}


