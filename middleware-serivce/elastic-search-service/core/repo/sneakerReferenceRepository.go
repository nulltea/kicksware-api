package repo

import model "elastic-search-service/core/model"


type SneakerReferenceRepository interface {
	FetchOne(uniqueId string) (*model.SneakerReference, error)
	Fetch(uniqueId []string) ([]*model.SneakerReference, error)
	FetchAll() ([]*model.SneakerReference, error)
	FetchQuery(query interface{}) ([]*model.SneakerReference, error)
	Count(query interface{}) (int64, error)
}
