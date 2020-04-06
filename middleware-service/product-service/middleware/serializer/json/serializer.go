package json

import (
	"encoding/json"

	"github.com/pkg/errors"

	"product-service/core/model"
	"product-service/core/service"
)

type serializer struct{}

func NewSerializer() service.SneakerProductSerializer {
	return &serializer{}
}

func (r *serializer) Decode(input []byte) (*model.SneakerProduct, error) {
	sneakerProduct := &model.SneakerProduct{}
	if err := json.Unmarshal(input, sneakerProduct); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerProduct.Decode")
	}
	return sneakerProduct, nil
}

func (r *serializer) Encode(input interface{}) ([]byte, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerProduct.Encode")
	}
	return raw, nil
}