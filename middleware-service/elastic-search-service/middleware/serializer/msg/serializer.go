package msg

import (
	"elastic-search-service/core/model"
	"elastic-search-service/core/service"
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

type serializer struct{}

func NewSerializer() service.SneakerReferenceSerializer {
	return &serializer{}
}

func (r *serializer) Decode(input []byte) (*model.SneakerReference, error) {
	sneakerProduct := &model.SneakerReference{}
	if err := msgpack.Unmarshal(input, sneakerProduct); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.Decode")
	}
	return sneakerProduct, nil
}

func (r *serializer) Encode(input interface{}) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.Encode")
	}
	return rawMsg, nil
}
