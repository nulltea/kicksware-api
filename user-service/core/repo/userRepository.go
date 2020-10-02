package repo

import (
	"go.kicksware.com/api/service-common/core/meta"

	"go.kicksware.com/api/user-service/core/model"
)

type UserRepository interface {
	FetchOne(code string) (*model.User, error)
	Fetch(code []string, params *meta.RequestParams) ([]*model.User, error)
	FetchAll(params *meta.RequestParams) ([]*model.User, error)
	FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.User, error)
	Store(user *model.User) error
	Modify(user *model.User) error
	Replace(user *model.User) error
	Remove(code string) error
	Count(query meta.RequestQuery, params *meta.RequestParams) (int, error)
	CountAll() (int, error)
}