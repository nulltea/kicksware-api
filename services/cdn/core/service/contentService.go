package service

import (
	"io"

	"go.kicksware.com/api/services/cdn/core/meta"
	"go.kicksware.com/api/services/cdn/core/model"
)

type ContentService interface {
	Original(query meta.ContentQuery) (*model.Content, error)
	Crop(query meta.ContentQuery, options meta.ImageOptions) (*model.Content, error)
	Resize(query meta.ContentQuery, options meta.ImageOptions) (*model.Content, error)
	Thumbnail(query meta.ContentQuery) (*model.Content, error)
	Upload(r io.Reader, query meta.ContentQuery) error
}
