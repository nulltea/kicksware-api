package service

import "user-service/core/model"

type UserService interface {
	FetchOne(code string) (*model.User, error)
	Fetch(code []string) ([]*model.User, error)
	FetchAll() ([]*model.User, error)
	FetchQuery(query interface{}) ([]*model.User, error)
	Register(user *model.User) error
	Modify(user *model.User) error
	Replace(user *model.User) error
	Remove(code string) error
	RemoveObj(user *model.User) error
}