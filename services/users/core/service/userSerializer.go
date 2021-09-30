package service

import "go.kicksware.com/api/services/users/core/model"

type UserSerializer interface {
	Decode(input []byte) (*model.User, error)
	DecodeRange(input []byte) ([]*model.User, error)
	DecodeMap(input []byte) (map[string]interface{}, error)
	DecodeInto(input []byte, target interface{}) error
	Encode(input interface{}) ([]byte, error)
}
