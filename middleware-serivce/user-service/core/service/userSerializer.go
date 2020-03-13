package service

import model "user-service/core/model"

type UserSerializer interface {
	Decode(input []byte) (*model.User, error)
	Encode(input *model.User) ([]byte, error)
}
