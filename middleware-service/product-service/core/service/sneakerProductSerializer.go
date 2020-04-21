package service

import "product-service/core/model"

type SneakerProductSerializer interface {
	DecodeMap(input []byte) (map[string]interface{}, error)
	Decode(input []byte) (*model.SneakerProduct, error)
	Encode(input interface{}) ([]byte, error)
}
