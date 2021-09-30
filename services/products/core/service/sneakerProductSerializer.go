package service

import "go.kicksware.com/api/services/products/core/model"

type SneakerProductSerializer interface {
	Decode(input []byte) (*model.SneakerProduct, error)
	DecodeRange(input []byte) ([]*model.SneakerProduct, error)
	DecodeMap(input []byte) (map[string]interface{}, error)
	DecodeInto(input []byte, target interface{}) error
	Encode(input interface{}) ([]byte, error)
}
