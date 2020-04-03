package service

import model "elastic-search-service/core/model"

type SneakerReferenceSerializer interface {
	Decode(input []byte) (*model.SneakerReference, error)
	Encode(input interface{}) ([]byte, error)
}
