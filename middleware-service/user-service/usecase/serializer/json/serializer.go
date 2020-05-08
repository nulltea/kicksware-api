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

func (r *serializer) Decode(input []byte) (user *model.User, err error) {
	if err = json.Unmarshal(input, &user); err != nil {
		return nil, errors.Wrap(err, "serializer.User.Decode")
	}
	return
}

func (r *serializer) DecodeRange(input []byte) (users []*model.User, err error) {
	if err := json.Unmarshal(input, &users); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.DecodeRange")
	}
	return
}

func (r *serializer) DecodeMap(input []byte) (map[string]interface{}, error) {
	queryMap := make(map[string]interface{})
	if err := json.Unmarshal(input, &queryMap); err != nil {
		return nil, errors.Wrap(err, "serializer.User.DecodeMap")
	}
	return queryMap, nil
}

func (r *serializer) DecodeInto(input []byte, target interface{}) error  {
	if err := json.Unmarshal(input, target); err != nil {
		return errors.Wrap(err, "serializer.User.DecodeInto")
	}
	return nil
}

func (r *serializer) Encode(input interface{}) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.User.Encode")
	}
	return rawMsg, nil
}