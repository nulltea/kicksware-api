package service

import model "product-service/core/model"

type SneakerProductSerializer interface {
	Decode(input []byte) (*model.SneakerProduct, error)
	Encode(input *model.SneakerProduct) ([]byte, error)
}
