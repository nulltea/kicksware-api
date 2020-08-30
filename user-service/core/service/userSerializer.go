package service

import "github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/model"

type UserSerializer interface {
	Decode(input []byte) (*model.User, error)
	DecodeRange(input []byte) ([]*model.User, error)
	DecodeMap(input []byte) (map[string]interface{}, error)
	DecodeInto(input []byte, target interface{}) error
	Encode(input interface{}) ([]byte, error)
}
