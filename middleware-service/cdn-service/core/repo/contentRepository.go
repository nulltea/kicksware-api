package repo

import "os"

type ContentRepository interface {
	Download(from, filename string) (*os.File, error)
	Upload(to string, filename string, content []byte) (string, error)
}