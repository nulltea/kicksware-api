package service

import "reference-service/core/model"

type SneakerReferenceService interface {
	FetchOne(code string) (*model.SneakerReference, error)
	Fetch(codes []string) ([]*model.SneakerReference, error)
	FetchAll() ([]*model.SneakerReference, error)
	FetchQuery(query interface{}) ([]*model.SneakerReference, error)
	StoreOne(sneakerReference *model.SneakerReference) error
	Store(sneakerReference []*model.SneakerReference) error
	Modify(sneakerReferences *model.SneakerReference) error
}