package redis

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
	"go.kicksware.com/api/service-common/config"

	"go.kicksware.com/api/service-common/core/meta"

	"go.kicksware.com/api/product-service/core/model"
	"go.kicksware.com/api/product-service/core/repo"
	"go.kicksware.com/api/product-service/usecase/business"
)

type repository struct {
	client *redis.Client
}

func newRedisClient(redisURL string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewRedisRepository(config config.DataStoreConfig) (repo.SneakerProductRepository, error) {
	rep := &repository{}
	client, err := newRedisClient(config.URL)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRedisRepository")
	}
	rep.client = client
	return rep, nil
}

func (r *repository) generateKey(code  string) string {
	return fmt.Sprintf("sneakerProduct[%s]", code)
}

func (r *repository) FetchOne(code string) (*model.SneakerProduct, error) {
	key := r.generateKey(code)
	data, err := r.client.Get(key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.Fetch")
	}
	if data == nil {
		return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.Fetch")
	}
	sneakerProduct := &model.SneakerProduct{}
	if err = json.Unmarshal(data, sneakerProduct); err != nil{
		return nil, err
	}
	return sneakerProduct, nil
}

func (r *repository) Fetch(codes []string, params *meta.RequestParams) ([]*model.SneakerProduct, error) {
	keys := funk.Map(codes, r.generateKey).([]string)
	data, err := r.client.MGet(keys...).Result()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.Fetch")
	}
	if len(data) == 0 {
		return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.Fetch")
	}
	sneakerProducts := funk.Map(data, func(val interface{}) (s *model.SneakerProduct) {
		json.Unmarshal([]byte(val.(string)), s)
		return
	} ).([]*model.SneakerProduct)
	return sneakerProducts, nil
}

func (r *repository) FetchAll(params *meta.RequestParams) ([]*model.SneakerProduct, error) {
	keys := r.client.Keys("sneakerProduct*").Val()
	if len(keys) == 0 {
		return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.FetchAll")
	}
	data, err := r.client.MGet(keys...).Result()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.FetchAll")
	}
	if len(data) == 0 {
		return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.FetchAll")
	}
	sneakerProducts := funk.Map(data, func(val interface{}) (s *model.SneakerProduct) {
		json.Unmarshal([]byte(val.(string)), s)
		return
	} ).([]*model.SneakerProduct)
	return sneakerProducts, nil
}

func (r *repository) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.SneakerProduct, error) {
	return r.FetchAll(params) //todo querying
}

func (r *repository) Store(sneakerProduct *model.SneakerProduct) error {
	key := r.generateKey(sneakerProduct.UniqueId)
	data, err := json.Marshal(sneakerProduct)
	if err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Store")
	}
	if err = r.client.MSet(key, data).Err(); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Store")
	}
	return nil
}

func (r *repository) Modify(sneakerProduct *model.SneakerProduct) error {
	if err := r.Store(sneakerProduct); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Modify")
	}
	return nil
}

func (r *repository) Replace(sneakerProduct *model.SneakerProduct) error {
	if err := r.Store(sneakerProduct); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Replace")
	}
	return nil
}

func (r *repository) Remove(code string) error {
	key := r.generateKey(code)
	if err := r.client.Del(key).Err(); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Remove")
	}
	return nil
}

func (r *repository) Count(query meta.RequestQuery, params *meta.RequestParams) (int, error) {
	return r.CountAll()
}

func (r *repository) CountAll() (count int, err error) {
	keys := r.client.Keys("sneakerProduct*").Val()
	return int(len(keys)), nil
}

