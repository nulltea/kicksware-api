package msg

import (
	"go.kicksware.com/api/beta/core/model"
	"go.kicksware.com/api/beta/core/service"

	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

type serializer struct{}

func NewSerializer() service.BetaSerializer {
	return &serializer{}
}

func (r *serializer) Decode(input []byte) (ref *model.Beta, err error) {
	if err := msgpack.Unmarshal(input, &ref); err != nil {
		return nil, errors.Wrap(err, "serializer.Beta.Decode")
	}
	return
}

func (r *serializer) DecodeRange(input []byte) (refs []*model.Beta, err error) {
	if err := msgpack.Unmarshal(input, &refs); err != nil {
		return nil, errors.Wrap(err, "serializer.Beta.DecodeRange")
	}
	return
}

func (r *serializer) DecodeMap(input []byte) (map[string]interface{}, error) {
	queryMap := make(map[string]interface{})
	if err := msgpack.Unmarshal(input, &queryMap); err != nil {
		return nil, errors.Wrap(err, "serializer.Beta.DecodeRange")
	}
	return queryMap, nil
}

func (r *serializer) DecodeInto(input []byte, target interface{}) error  {
	if err := msgpack.Unmarshal(input, target); err != nil {
		return errors.Wrap(err, "serializer.Beta.DecodeInto")
	}
	return nil
}

func (r *serializer) Encode(input interface{}) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Beta.Encode")
	}
	return rawMsg, nil
}
