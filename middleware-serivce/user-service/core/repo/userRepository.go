package repo

import model "user-service/core/model"

type UserRepository interface {
	RetrieveOne(uniqueId string) (*model.User, error)
	Retrieve(uniqueId []string) ([]*model.User, error)
	RetrieveAll() ([]*model.User, error)
	RetrieveQuery(query interface{}) ([]*model.User, error)
	Store(user *model.User) error
	Modify(user *model.User) error
	Remove(user *model.User) error
}