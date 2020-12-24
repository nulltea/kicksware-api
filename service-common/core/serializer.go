package core

type Serializer interface {
	Encode(data interface{}) ([]byte, error)
}