package pipe

import "github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/model"

type SneakerProductPipe interface {
	FetchOne(uniqueId string) (*model.SneakerProduct, error)
	Fetch(uniqueId []string) ([]*model.SneakerProduct, error)
	FetchAll() ([]*model.SneakerProduct, error)
	FetchQuery(query interface{}) ([]*model.SneakerProduct, error)
}