package json

import (
	"encoding/json"
	"github.com/pkg/errors"
	"product-service/core/model"
)

type SneakerProduct struct{}

func (r *SneakerProduct) Decode(input []byte) (*model.SneakerProduct, error) {
	sneakerProduct := &model.SneakerProduct{}
	if err := json.Unmarshal(input, sneakerProduct); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerProduct.Decode")
	}
	return sneakerProduct, nil
}

func (r *SneakerProduct) Encode(input *model.SneakerProduct) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerProduct.Encode")
	}
	return rawMsg, nil
}