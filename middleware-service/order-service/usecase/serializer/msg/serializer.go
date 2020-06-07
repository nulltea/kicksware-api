package msg

import (
	"order-service/core/model"
	"order-service/core/service"

	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

type serializer struct{}

func NewSerializer() service.OrderSerializer {
	return &serializer{}
}

func (r *serializer) Decode(input []byte) (ref *model.Order, err error) {
	if err := msgpack.Unmarshal(input, &ref); err != nil {
		return nil, errors.Wrap(err, "serializer.Order.Decode")
	}
	return
}

func (r *serializer) DecodeRange(input []byte) (refs []*model.Order, err error) {
	if err := msgpack.Unmarshal(input, &refs); err != nil {
		return nil, errors.Wrap(err, "serializer.Order.DecodeRange")
	}
	return
}

func (r *serializer) DecodeMap(input []byte) (map[string]interface{}, error) {
	queryMap := make(map[string]interface{})
	if err := msgpack.Unmarshal(input, &queryMap); err != nil {
		return nil, errors.Wrap(err, "serializer.Order.DecodeRange")
	}
	return queryMap, nil
}

func (r *serializer) DecodeInto(input []byte, target interface{}) error  {
	if err := msgpack.Unmarshal(input, target); err != nil {
		return errors.Wrap(err, "serializer.Order.DecodeInto")
	}
	return nil
}

func (r *serializer) Encode(input interface{}) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Order.Encode")
	}
	return rawMsg, nil
}
