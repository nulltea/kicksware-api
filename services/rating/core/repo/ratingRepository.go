package repo

import (
	"go.kicksware.com/api/shared/core/meta"

	"go.kicksware.com/api/services/rating/core/model"
)


type RatingRepository interface {
	FetchOne(code string, params *meta.RequestParams) (*model.Rating, error)
	Fetch(codes []string, params *meta.RequestParams) ([]*model.Rating, error)
	FetchAll(params *meta.RequestParams) ([]*model.Rating, error)
	FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.Rating, error)
	StoreOne(rate *model.Rating) error
	Modify(rates *model.Rating) error
	Remove(code string) error
	Count(query meta.RequestQuery, params *meta.RequestParams) (int, error)
	CountAll() (int, error)
}
