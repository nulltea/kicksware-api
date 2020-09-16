package service

import (
	ref "github.com/timoth-y/kicksware-api/reference-service/core/model"
	prod "github.com/timoth-y/kicksware-api/product-service/core/model"
)

type SneakerSearchSerializer interface {
	DecodeReference(input []byte) (*ref.SneakerReference, error)
	DecodeReferences(input []byte) ([]*ref.SneakerReference, error)
	DecodeProduct(input []byte) (*prod.SneakerProduct, error)
	DecodeProducts(input []byte) ([]*prod.SneakerProduct, error)
	DecodeMap(input []byte) (map[string]interface{}, error)
	DecodeInto(input []byte, target interface{}) error
	Encode(input interface{}) ([]byte, error)
}
