package json

import (
	"encoding/json"

	"github.com/pkg/errors"

	"go.kicksware.com/api/services/rating/core/model"
	"go.kicksware.com/api/services/rating/core/service"
)

type serializer struct{}

func NewSerializer() service.RatingSerializer {
	return &serializer{}
}

func (r *serializer) Decode(input []byte) (ref *model.Rating, err error) {
	if err = json.Unmarshal(input, &ref); err != nil {
		return nil, errors.Wrap(err, "serializer.Rating.Decode")
	}
	return
}

func (r *serializer) DecodeRange(input []byte) (refs []*model.Rating, err error) {
	if err = json.Unmarshal(input, &refs); err != nil {
		return nil, errors.Wrap(err, "serializer.Rating.DecodeRange")
	}
	return
}

func (r *serializer) DecodeMap(input []byte) (map[string]interface{}, error) {
	queryMap := make(map[string]interface{})
	if err := json.Unmarshal(input, &queryMap); err != nil {
		return nil, errors.Wrap(err, "serializer.Rating.DecodeRange")
	}
	return queryMap, nil
}

func (r *serializer) DecodeInto(input []byte, target interface{}) error  {
	if err := json.Unmarshal(input, target); err != nil {
		return errors.Wrap(err, "serializer.Rating.DecodeInto")
	}
	return nil
}

func (r *serializer) Encode(input interface{}) ([]byte, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Rating.Encode")
	}
	return raw, nil
}
