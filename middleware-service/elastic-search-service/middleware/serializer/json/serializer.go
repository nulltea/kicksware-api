package json

import (
	"encoding/json"
	"github.com/pkg/errors"
	"elastic-search-service/core/model"
	"elastic-search-service/core/service"
)

type serializer struct{}

func NewSerializer() service.SneakerReferenceSerializer {
	return &serializer{}
}

func (r *serializer) Decode(input []byte) (*model.SneakerReference, error) {
	sneakerProduct := &model.SneakerReference{}
	if err := json.Unmarshal(input, sneakerProduct); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.Decode")
	}
	return sneakerProduct, nil
}

func (r *serializer) Encode(input interface{}) ([]byte, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.Encode")
	}
	return raw, nil
}