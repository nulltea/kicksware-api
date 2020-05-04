package service

import (
	"user-service/core/meta"
	"user-service/core/model"
)

type UserService interface {
	FetchOne(code string) (*model.User, error)
	Fetch(code []string) ([]*model.User, error)
	FetchAll() ([]*model.User, error)
	FetchQuery(query meta.RequestQuery) ([]*model.User, error)
	Register(user *model.User) error
	Modify(user *model.User) error
	Replace(user *model.User) error
	Remove(code string) error
	Count(query meta.RequestQuery) (int, error)
	CountAll() (int, error)
}