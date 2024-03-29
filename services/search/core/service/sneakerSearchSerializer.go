package service

import (
	ref "go.kicksware.com/api/services/references/core/model"
	prod "go.kicksware.com/api/services/products/core/model"
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
