package msg

import (
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
	"user-service/core/model"
)

type User struct{}

func (r *User) Decode(input []byte) (*model.User, error) {
	user := &model.User{}
	if err := msgpack.Unmarshal(input, user); err != nil {
		return nil, errors.Wrap(err, "serializer.User.Decode")
	}
	return user, nil
}

func (r *User) Encode(input *model.User) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.User.Encode")
	}
	return rawMsg, nil
}
