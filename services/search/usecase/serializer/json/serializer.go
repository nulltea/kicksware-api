package json

import (
	"encoding/json"

	"github.com/pkg/errors"

	prod "go.kicksware.com/api/services/products/core/model"
	ref "go.kicksware.com/api/services/references/core/model"

	"go.kicksware.com/api/services/search/core/service"
)

type serializer struct{}

func NewSerializer() service.SneakerSearchSerializer {
	return &serializer{}
}

func (r *serializer) DecodeReference(input []byte) (ref *ref.SneakerReference, err error) {
	if err = json.Unmarshal(input, &ref); err != nil {
		return nil, errors.Wrap(err, "serializer.Search.DecodeReference")
	}
	return
}

func (r *serializer) DecodeProduct(input []byte) (prod *prod.SneakerProduct, err error) {
	if err = json.Unmarshal(input, &prod); err != nil {
		return nil, errors.Wrap(err, "serializer.Search.DecodeProduct")
	}
	return
}

func (r *serializer) DecodeReferences(input []byte) (refs []*ref.SneakerReference, err error) {
	if err = json.Unmarshal(input, &refs); err != nil {
		return nil, errors.Wrap(err, "serializer.Search.DecodeReferences")
	}
	return
}

func (r *serializer) DecodeProducts(input []byte) (products []*prod.SneakerProduct, err error) {
	if err = json.Unmarshal(input, &products); err != nil {
		return nil, errors.Wrap(err, "serializer.Search.DecodeProducts")
	}
	return
}

func (r *serializer) DecodeMap(input []byte) (map[string]interface{}, error) {
	queryMap := make(map[string]interface{})
	if err := json.Unmarshal(input, &queryMap); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.DecodeMap")
	}
	return queryMap, nil
}

func (r *serializer) DecodeInto(input []byte, target interface{}) error  {
	if err := json.Unmarshal(input, target); err != nil {
		return errors.Wrap(err, "serializer.SneakerReference.DecodeInto")
	}
	return nil
}

func (r *serializer) Encode(input interface{}) ([]byte, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.Encode")
	}
	return raw, nil
}
