package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"middleware-serivce/model"
	"net/http"
	"strconv"
)

type redisRepository struct {
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

func NewRedisRepository(redisURL string) (model.SneakerProductRepository, error) {
	repo := &redisRepository{}
	client, err := newRedisClient(redisURL)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRedisRepository")
	}
	repo.client = client
	return repo, nil
}

func (r *redisRepository) generateKey(id string) string {
	return fmt.Sprintf("sneakerProduct:%s", id)
}

func (r *redisRepository) Find(id string) (*model.SneakerProduct, error) {
	sneakerProduct := &model.SneakerProduct{}
	key := r.generateKey(id)
	data, err := r.client.HGetAll(key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.Find")
	}
	if len(data) == 0 {
		return nil, errors.Wrap(model.ErrProductNotFound, "repository.SneakerProduct.Find")
	}
	createdAt, err := http.ParseTime(data["created_at"])
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.Find")
	}
	sneakerProduct.Id = data["Id"]
	sneakerProduct.BrandName = data["BrandName"]
	sneakerProduct.ModelName = data["ModelName"]
	sneakerProduct.Owner = data["Owner"]
	sneakerProduct.URL = data["URL"]
	if stateIndex, err := strconv.ParseFloat(data["StateIndex"], 32); err == nil {
		sneakerProduct.StateIndex = float32(stateIndex)
	}
	sneakerProduct.AddedAt = createdAt
	return sneakerProduct, nil
}

func (r *redisRepository) Store(sneakerProduct *model.SneakerProduct) error {
	key := r.generateKey(sneakerProduct.Id)
	data := map[string]interface{} {
		"Id":        sneakerProduct.Id,
		"ModelName": sneakerProduct.BrandName,
		"BrandName": sneakerProduct.ModelName,
		"Owner":     sneakerProduct.Owner,
		"StateIndex": sneakerProduct.StateIndex,
		"URL":       sneakerProduct.URL,
		"AddedAt":   sneakerProduct.AddedAt,
	}
	_, err := r.client.HMSet(key, data).Result()
	if err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Store")
	}
	return nil
}
