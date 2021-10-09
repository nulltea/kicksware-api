package business

import (
	"bytes"
	"image"
	"io"

	"github.com/disintegration/imaging"
	"github.com/golang/glog"
	"go.kicksware.com/api/shared/config"

	"go.kicksware.com/api/services/cdn/core/meta"
	"go.kicksware.com/api/services/cdn/core/model"
	"go.kicksware.com/api/services/cdn/core/repo"
	"go.kicksware.com/api/services/cdn/core/service"
)

type contentService struct {
	repo          repo.ContentRepository
	serviceConfig config.CommonConfig
}

func NewContentService(repo repo.ContentRepository, config config.CommonConfig) service.ContentService {
	return &contentService{
		repo,
		config,
	}
}

func (s *contentService) Original(query meta.ContentQuery) (*model.Content, error) {
	img, mimeType, err := s.downloadImage(query); if err != nil {
		glog.Errorln(err)
		return nil, err
	}
	format, err := imaging.FormatFromFilename(query.Filename); if err != nil {
		glog.Errorln(err)
		return nil, err
	}
	buffer := &bytes.Buffer{}
	imaging.Encode(buffer, img, format)
	return imageContentOf(buffer.Bytes(), mimeType), nil
}

func (s *contentService) Crop(query meta.ContentQuery, options meta.ImageOptions) (*model.Content, error) {
	img, mimeType, err := s.downloadImage(query); if err != nil {
		glog.Errorln(err)
		return nil, err
	}
	format, err := imaging.FormatFromFilename(query.Filename); if err != nil {
		glog.Errorln(err)
		return nil, err
	}
	img = imaging.CropCenter(img, options.Width, options.Height)
	buffer := &bytes.Buffer{}
	imaging.Encode(buffer, img, format)
	return imageContentOf(buffer.Bytes(), mimeType), nil
}

func (s *contentService) Resize(query meta.ContentQuery, options meta.ImageOptions) (*model.Content, error) {
	img, mimeType, err := s.downloadImage(query); if err != nil {
		glog.Errorln(err)
		return nil, err
	}
	format, err := imaging.FormatFromFilename(query.Filename); if err != nil {
		glog.Errorln(err)
		return nil, err
	}
	img = imaging.Resize(img, options.Width, options.Height, imaging.Lanczos)
	buffer := &bytes.Buffer{}
	imaging.Encode(buffer, img, format)
	return imageContentOf(buffer.Bytes(), mimeType), nil
}

func (s *contentService) Thumbnail(query meta.ContentQuery) (*model.Content, error) {
	img,  mimeType, err := s.downloadImage(query); if err != nil {
		glog.Errorln(err)
		return nil, err
	}
	format, err := imaging.FormatFromFilename(query.Filename); if err != nil {
		glog.Errorln(err)
		return nil, err
	}
	img = imaging.Thumbnail(img, 100, 100, imaging.Lanczos)
	buffer := &bytes.Buffer{}
	imaging.Encode(buffer, img, format)
	return imageContentOf(buffer.Bytes(), mimeType), nil
}

func (s *contentService) Upload(r io.Reader, query meta.ContentQuery) error {
	panic("implement me")
}

func (s *contentService) downloadImage(query meta.ContentQuery) (image.Image, meta.MimeType, error) {
	file, err := s.repo.Download(query.Collection, query.Filename)
	if err != nil {
		return nil, "", err
	}

	img, mimeType, err := image.Decode(bytes.NewBuffer(file)); if err != nil {
		glog.Error(err)
		return nil, "", err
	}

	if err != nil {
		glog.Errorln(err)
		return  nil, "", err
	}
	return img, mimeTypeOf(mimeType), nil
}

func imageContentOf(data []byte, mimeType meta.MimeType) *model.Content {
	return &model.Content{
		data,
		mimeType,
	}
}

func setMaxSize(max int, size []int) []int {
	if max <= size[0] {
		size[0] = max
	}
	if max <= size[1] {
		size[1] = max
	}
	return size
}

func fitToActualSize(img *image.Image, size []int) []int {
	ib := (*img).Bounds()
	var x, y int = ib.Dx(), ib.Dy()
	if ib.Dx() >= size[0] {
		x = size[0]
	}
	if ib.Dy() >= size[1] {
		y = size[1]
	}
	return []int{x, y}
}

func mimeTypeOf(imgType string) meta.MimeType {
	switch imgType {
	case "png":
		return "image/png"
	case "webp":
		return "image/webp"
	case "tiff":
		return "image/tiff"
	case "gif":
		return "image/gif"
	case "svg":
		return "image/svg+xml"
	case "pdf":
		return "application/pdf"
	default:
		return "image/jpeg"
	}
}
