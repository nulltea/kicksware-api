package json

import (
	"encoding/json"
	"github.com/pkg/errors"
	"user-service/core/model"
	"user-service/core/service"
)

type serializer struct{}

func NewSerializer() service.UserSerializer {
	return &serializer{}
}


func (r *serializer) Decode(input []byte) (*model.User, error) {
	user := &model.User{}
	if err := json.Unmarshal(input, user); err != nil {
		return nil, errors.Wrap(err, "serializer.User.Decode")
	}
	return user, nil
}

func (r *serializer) Encode(input interface{}) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.User.Encode")
	}
	return rawMsg, nil
}