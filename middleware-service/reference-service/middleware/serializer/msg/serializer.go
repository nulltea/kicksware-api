package msg

import (
	"reference-service/core/model"
	"reference-service/core/service"

	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

type serializer struct{}

func NewSerializer() service.SneakerReferenceSerializer {
	return &serializer{}
}

func (r *serializer) DecodeOne(input []byte) (*model.SneakerReference, error) {
	reference := &model.SneakerReference{}
	if err := msgpack.Unmarshal(input, reference); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.DecodeOne")
	}
	return reference, nil
}

func (r *serializer) Decode(input []byte) ([]*model.SneakerReference, error) {
	references := make([]*model.SneakerReference, 0)
	if err := msgpack.Unmarshal(input, references); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.Decode")
	}
	return references, nil
}

func (r *serializer) DecodeMap(input []byte) (map[string]interface{}, error) {
	queryMap := make(map[string]interface{})
	if err := msgpack.Unmarshal(input, &queryMap); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.Decode")
	}
	return queryMap, nil
}

func (r *serializer) Encode(input interface{}) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.Encode")
	}
	return rawMsg, nil
}
