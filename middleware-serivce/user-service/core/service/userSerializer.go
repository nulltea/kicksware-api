package service

import "user-service/core/model"

type UserSerializer interface {
	Decode(input []byte) (*model.User, error)
	Encode(input interface{}) ([]byte, error)
}
