package service

import "product-service/core/model"

type SneakerProductSerializer interface {
	Decode(input []byte) (*model.SneakerProduct, error)
	Encode(input interface{}) ([]byte, error)
}
