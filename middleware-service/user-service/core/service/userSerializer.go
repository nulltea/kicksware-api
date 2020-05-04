package service

import "user-service/core/model"

type UserSerializer interface {
	Decode(input []byte) (*model.User, error)
	Encode(input interface{}) ([]byte, error)
	DecodeMap(input []byte) (map[string]interface{}, error)
	DecodeInto(input []byte, target interface{}) error
}
