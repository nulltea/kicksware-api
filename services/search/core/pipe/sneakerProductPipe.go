package pipe

import "go.kicksware.com/api/services/products/core/model"

type SneakerProductPipe interface {
	FetchOne(uniqueId string) (*model.SneakerProduct, error)
	Fetch(uniqueId []string) ([]*model.SneakerProduct, error)
	FetchAll() ([]*model.SneakerProduct, error)
	FetchQuery(query interface{}) ([]*model.SneakerProduct, error)
}
