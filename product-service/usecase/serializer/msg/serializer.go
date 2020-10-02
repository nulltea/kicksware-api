package msg

import (
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"

	"go.kicksware.com/api/product-service/core/model"
	"go.kicksware.com/api/product-service/core/service"
)

type serializer struct{}

func NewSerializer() service.SneakerProductSerializer {
	return &serializer{}
}

func (r *serializer) Decode(input []byte) (product *model.SneakerProduct, err error) {
	if err = msgpack.Unmarshal(input, &product); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerProduct.Decode")
	}
	return
}

func (r *serializer) DecodeRange(input []byte) (products []*model.SneakerProduct, err error) {
	if err = msgpack.Unmarshal(input, &products); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerProduct.DecodeRange")
	}
	return
}

func (r *serializer) DecodeMap(input []byte) (map[string]interface{}, error) {
	queryMap := make(map[string]interface{})
	if err := msgpack.Unmarshal(input, &queryMap); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerProduct.Decode")
	}
	return queryMap, nil
}

func (r *serializer) DecodeInto(input []byte, target interface{}) error  {
	if err := msgpack.Unmarshal(input, target); err != nil {
		return errors.Wrap(err, "serializer.SneakerProduct.DecodeInto")
	}
	return nil
}

func (r *serializer) Encode(input interface{}) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerProduct.Encode")
	}
	return rawMsg, nil
}
