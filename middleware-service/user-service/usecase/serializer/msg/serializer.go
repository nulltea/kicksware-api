package msg

import (
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"

	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/model"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/service"
)

type serializer struct{}

func NewSerializer() service.UserSerializer {
	return &serializer{}
}

func (r *serializer) Decode(input []byte) (user *model.User, err error) {
	if err = msgpack.Unmarshal(input, &user); err != nil {
		return nil, errors.Wrap(err, "serializer.User.Decode")
	}
	return user, nil
}

func (r *serializer) DecodeRange(input []byte) (users []*model.User, err error) {
	if err := msgpack.Unmarshal(input, &users); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.DecodeRange")
	}
	return
}

func (r *serializer) DecodeMap(input []byte) (map[string]interface{}, error) {
	queryMap := make(map[string]interface{})
	if err := msgpack.Unmarshal(input, &queryMap); err != nil {
		return nil, errors.Wrap(err, "serializer.User.Decode")
	}
	return queryMap, nil
}

func (r *serializer) DecodeInto(input []byte, target interface{}) error  {
	if err := msgpack.Unmarshal(input, target); err != nil {
		return errors.Wrap(err, "serializer.User.DecodeInto")
	}
	return nil
}

func (r *serializer) Encode(input interface{}) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.User.Encode")
	}
	return rawMsg, nil
}
