package model

import (
	"io"

	"go.kicksware.com/api/cdn-service/core/meta"
)

type Content struct {
	Data []byte
	MimeType meta.MimeType
}

func (m *Content) Write(w io.Writer) (int, error){
	return w.Write(m.Data)
}