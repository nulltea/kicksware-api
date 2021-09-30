package proto

import (
	"go.kicksware.com/api/services/cdn/core/meta"
	"go.kicksware.com/api/services/cdn/core/model"
)

func (m *Content) ToNative() *model.Content {
	return &model.Content{
		Data:     m.Data,
		MimeType: meta.MimeType(m.MimeType),
	}
}

func (m Content) FromNative(n *model.Content) *Content {
	m.Data = n.Data
	m.MimeType = string(n.MimeType)
	return &m
}

func (m ContentRequest) ToNative() meta.ContentQuery {
	return meta.ContentQuery{
		Collection: m.Collection,
		Filename: m.Filename,
		ImageOptions: m.ImageOptions.ToNative(),
	}
}

func (m ContentRequest) FromNative(n meta.ContentQuery) ContentRequest {
	m.Collection = n.Collection
	m.Filename = n.Filename
	m.ImageOptions = ImageOptions{}.FromNative(n.ImageOptions)
	return m
}

func (m ImageOptions) ToNative() meta.ImageOptions {
	return meta.ImageOptions{
		Width: int(m.Width),
		Height: int(m.Height),
	}
}

func (m ImageOptions) FromNative(n meta.ImageOptions) *ImageOptions {
	m.Width = int64(n.Width)
	m.Height = int64(n.Height)
	return &m
}
