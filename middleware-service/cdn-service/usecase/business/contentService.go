package business

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"cdn-service/core/repo"
	"cdn-service/core/service"
	"cdn-service/env"
)

type contentService struct {
	repo repo.ContentRepository
	serviceConfig env.CommonConfig
}

func NewContentService(repo repo.ContentRepository, config env.CommonConfig) service.ContentService {
	return &contentService{
		repo,
		config,
	}
}

func (s *contentService) Original(w io.Writer, from, filename string) error {
	// file, err := s.repo.Download(from, filename)
	// writeByMimetype(w, )
	panic("implement me")
}

func (s *contentService) Crop(w io.Writer, from, filename string) error {
	panic("implement me")
}

func (s *contentService) Resize(w io.Writer, from, filename string) error {
	panic("implement me")
}

func (s *contentService) Upload(r io.Reader, to string) error {
	panic("implement me")
}

func writeByMimetype(w io.Writer, dst image.Image, mimetype string) error {
	switch mimetype {
	case "jpeg", "jpg":
		return jpeg.Encode(w, dst, &jpeg.Options{Quality: jpeg.DefaultQuality})
	case "png":
		return png.Encode(w, dst)
	case "svg":
		panic("not implemented")
	default:
		return fmt.Errorf("Mimetype '%s' can't be processed.", mimetype)
	}
}
