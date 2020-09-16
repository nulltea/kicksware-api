package service

import (
	"io"

	"github.com/timoth-y/kicksware-api/cdn-service/core/meta"
	"github.com/timoth-y/kicksware-api/cdn-service/core/model"
)

type ContentService interface {
	Original(query meta.ContentQuery) (*model.Content, error)
	Crop(query meta.ContentQuery, options meta.ImageOptions) (*model.Content, error)
	Resize(query meta.ContentQuery, options meta.ImageOptions) (*model.Content, error)
	Thumbnail(query meta.ContentQuery) (*model.Content, error)
	Upload(r io.Reader, query meta.ContentQuery) error
}