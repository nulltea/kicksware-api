package repo

import model "product-service/core/model"


type SneakerProductRepository interface {
	RetrieveOne(uniqueId string) (*model.SneakerProduct, error)
	Retrieve(uniqueId []string) ([]*model.SneakerProduct, error)
	RetrieveAll() ([]*model.SneakerProduct, error)
	RetrieveQuery(query interface{}) ([]*model.SneakerProduct, error)
	Store(sneakerProduct *model.SneakerProduct) error
	Modify(sneakerProduct *model.SneakerProduct) error
	Remove(sneakerProduct *model.SneakerProduct) error
}
