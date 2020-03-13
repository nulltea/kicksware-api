package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"user-service/core/model"
	"user-service/core/repo"
	"user-service/scenario/business"
)

type RedisRepository struct {
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
	repo := &RedisRepository{}
	client, err := newRedisClient(redisURL)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRedisRepository")
	}
	repo.client = client
	return repo, nil
}

func (r *RedisRepository) generateKey(id string) string {
	return fmt.Sprintf("sneakerProduct:%s", id)
}

func (r *RedisRepository) Retrieve(id string) (*model.SneakerProduct, error) {
	sneakerProduct := &model.SneakerProduct{}
	key := r.generateKey(id)
	data, err := r.client.HGetAll(key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.Retrieve")
	}
	if len(data) == 0 {
		return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.Retrieve")
	}
	createdAt, err := http.ParseTime(data["created_at"])
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.Retrieve")
	}
	sneakerProduct.UniqueId = data["UniqueId"]
	sneakerProduct.BrandName = data["BrandName"]
	sneakerProduct.ModelName = data["ModelName"]
	sneakerProduct.Owner = data["Owner"]
	if stateIndex, err := strconv.ParseFloat(data["ConditionIndex"], 64); err == nil {
		sneakerProduct.ConditionIndex = stateIndex
	}
	sneakerProduct.AddedAt = createdAt
	return sneakerProduct, nil
}

func (r *RedisRepository) Store(sneakerProduct *model.SneakerProduct) error {
	key := r.generateKey(sneakerProduct.UniqueId)
	data := map[string]interface{} {
		"UniqueId":        sneakerProduct.UniqueId,
		"ModelName": sneakerProduct.BrandName,
		"BrandName": sneakerProduct.ModelName,
		"Owner":     sneakerProduct.Owner,
		"ConditionIndex": sneakerProduct.ConditionIndex,
		"AddedAt":   sneakerProduct.AddedAt,
	}
	_, err := r.client.HMSet(key, data).Result()
	if err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Store")
	}
	return nil
}
