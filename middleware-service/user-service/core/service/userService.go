package service

import (
	"user-service/core/meta"
	"user-service/core/model"
)

type UserService interface {
	FetchOne(code string) (*model.User, error)
	Fetch(code []string, params meta.RequestParams) ([]*model.User, error)
	FetchByEmail(email string) (*model.User, error)
	FetchByUsername(username string) (*model.User, error)
	FetchAll(params meta.RequestParams) ([]*model.User, error)
	FetchQuery(query meta.RequestQuery, params meta.RequestParams) ([]*model.User, error)
	Register(user *model.User) error
	GenerateUsername(user *model.User, save bool) (string, error)
	Modify(user *model.User) error
	Replace(user *model.User) error
	Remove(code string) error
	Count(query meta.RequestQuery, params meta.RequestParams) (int, error)
	CountAll() (int, error)
}