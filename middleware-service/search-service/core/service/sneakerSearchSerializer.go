package service

import model "github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/model"

type SneakerSearchSerializer interface {
	DecodeReference(input []byte) (*model.SneakerReference, error)
	DecodeReferences(input []byte) ([]*model.SneakerReference, error)
	DecodeProduct(input []byte) (*model.SneakerProduct, error)
	DecodeProducts(input []byte) ([]*model.SneakerProduct, error)
	DecodeMap(input []byte) (map[string]interface{}, error)
	DecodeInto(input []byte, target interface{}) error
	Encode(input interface{}) ([]byte, error)
}
