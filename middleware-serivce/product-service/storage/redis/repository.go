package redis

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
	"product-service/core/model"
	"product-service/core/repo"
	"product-service/scenario/business"
)

type Repository struct {
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

func NewRedisRepository(redisURL string) (repo.SneakerProductRepository, error) {
	repo := &Repository{}
	client, err := newRedisClient(redisURL)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRedisRepository")
	}
	repo.client = client
	return repo, nil
}

func (r *Repository) generateKey(code  string) string {
	return fmt.Sprintf("sneakerProduct[%s]", id)
}

func (r *Repository) RetrieveOne(code string) (*model.SneakerProduct, error) {
	key := r.generateKey(code)
	data, err := r.client.Get(key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.Retrieve")
	}
	if data == nil {
		return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.Retrieve")
	}
	sneakerProduct := &model.SneakerProduct{}
	if err = json.Unmarshal(data, sneakerProduct); err != nil{
		return nil, err
	}
	return sneakerProduct, nil
}

func (r *Repository) Retrieve(codes []string) ([]*model.SneakerProduct, error) {
	keys := funk.Map(codes, r.generateKey).([]string)
	data, err := r.client.MGet(keys...).Result()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.Retrieve")
	}
	if len(data) == 0 {
		return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.Retrieve")
	}
	sneakerProducts := funk.Map(data, func(val interface{}) (s *model.SneakerProduct) {
		json.Unmarshal([]byte(val.(string)), s)
		return
	} ).([]*model.SneakerProduct)
	return sneakerProducts, nil
}

func (r *Repository) RetrieveAll() ([]*model.SneakerProduct, error) {
	keys := r.client.Keys("sneakerProduct*").Val()
	if len(keys) == 0 {
		return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.RetrieveAll")
	}
	data, err := r.client.MGet(keys...).Result()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.RetrieveAll")
	}
	if len(data) == 0 {
		return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.RetrieveAll")
	}
	sneakerProducts := funk.Map(data, func(val interface{}) (s *model.SneakerProduct) {
		json.Unmarshal([]byte(val.(string)), s)
		return
	} ).([]*model.SneakerProduct)
	return sneakerProducts, nil
}

func (r *Repository) RetrieveQuery(query interface{}) ([]*model.SneakerProduct, error) {
	return r.RetrieveAll() //todo querying
}

func (r *Repository) Store(sneakerProduct *model.SneakerProduct) error {
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

func (r *Repository) Modify(sneakerProduct *model.SneakerProduct) error {
	return r.Store(sneakerProduct)
}

func (r *Repository) Replace(sneakerProduct *model.SneakerProduct) error {
	return r.Store(sneakerProduct)
}

func (r *Repository) Remove(code string) error {
	key := r.generateKey(code)
	if err := r.client.Del(key).Err(); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Remove")
	}
	return nil
}

func (r *Repository) RemoveObj(sneakerProduct *model.SneakerProduct) error {
	return r.Remove(sneakerProduct.UniqueId)
}

func (r *Repository) Count(query interface{}) (int64, error) {
	keys := r.client.Keys("sneakerProduct*").Val()
	return int64(len(keys)), nil
}


