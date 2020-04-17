package json

import (
	"encoding/json"
	"github.com/pkg/errors"
	"reference-service/core/model"
	"reference-service/core/service"
)

type serializer struct{}

func NewSerializer() service.SneakerReferenceSerializer {
	return &serializer{}
}

func (r *serializer) DecodeOne(input []byte) (*model.SneakerReference, error) {
	reference := &model.SneakerReference{}
	if err := json.Unmarshal(input, reference); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.DecodeOne")
	}
	return reference, nil
}

func (r *serializer) Decode(input []byte) ([]*model.SneakerReference, error) {
	references := make([]*model.SneakerReference, 0)
	if err := json.Unmarshal(input, references); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.Decode")
	}
	return references, nil
}

func (r *serializer) Encode(input interface{}) ([]byte, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.Encode")
	}
	return raw, nil
}