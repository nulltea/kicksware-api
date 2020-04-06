package msg

import (
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"

	"product-service/core/model"
	"product-service/core/service"
)

type serializer struct{}

func NewSerializer() service.SneakerProductSerializer {
	return &serializer{}
}

func (r *serializer) Decode(input []byte) (*model.SneakerProduct, error) {
	sneakerProduct := &model.SneakerProduct{}
	if err := msgpack.Unmarshal(input, sneakerProduct); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerProduct.Decode")
	}
	return sneakerProduct, nil
}

func (r *serializer) Encode(input interface{}) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerProduct.Encode")
	}
	return rawMsg, nil
}
