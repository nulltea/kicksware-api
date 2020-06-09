package service

import "io"

type ContentService interface {
	Original(w io.Writer, from, filename string) error
	Crop(w io.Writer, from, filename string) error
	Resize(w io.Writer, from, filename string) error
	Upload(r io.Reader, to string) error
}