package repo

type ContentRepository interface {
	Download(from, filename string) ([]byte, error)
	Upload(to string, filename string, content []byte) (string, error)
}