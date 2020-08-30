package model

import (
	"io"

	"github.com/timoth-y/kicksware-platform/middleware-service/cdn-service/core/meta"
)

type Content struct {
	Data []byte
	MimeType meta.MimeType
}

func (m *Content) Write(w io.Writer) (int, error){
	return w.Write(m.Data)
}