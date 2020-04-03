package repo

import "user-service/core/model"

type UserRepository interface {
	FetchOne(code string) (*model.User, error)
	Fetch(code []string) ([]*model.User, error)
	FetchAll() ([]*model.User, error)
	FetchQuery(query interface{}) ([]*model.User, error)
	Store(user *model.User) error
	Modify(user *model.User) error
	Replace(user *model.User) error
	Remove(code string) error
	RemoveObj(user *model.User) error
}