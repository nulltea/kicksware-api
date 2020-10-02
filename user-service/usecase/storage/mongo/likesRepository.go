package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.kicksware.com/api/user-service/core/model"
	"go.kicksware.com/api/user-service/core/repo"
	"go.kicksware.com/api/user-service/env"
)

type likesRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	timeout    time.Duration
}

func NewLikesRepository(config env.DataStoreConfig) (repo.LikesRepository, error) {
	repo := &likesRepository{
		timeout: time.Duration(config.Timeout) * time.Second,
	}
	client, err := newMongoClient(config); if err != nil {
		return nil, errors.Wrap(err, "repository.NewLikesRepository")
	}
	repo.client = client
	database := client.Database(config.Database)
	repo.database = database
	repo.collection = database.Collection(config.LikesCollection)
	return repo, nil
}

func (r *likesRepository) AddLike(userID string, entityID string) error {
	like := &model.Like{
		UserID:   userID,
		EntityID: entityID,
	}
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, like)
	if err != nil {
		return errors.Wrap(err, "repository.Likes.AddLike")
	}
	return nil
}

func (r *likesRepository) RemoveLike(userID string, entityID string) error {
	like := &model.Like{
		UserID:   userID,
		EntityID: entityID,
	}
	filter, err := bson.Marshal(like); if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	_, err = r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "repository.Likes.RemoveLike")
	}
	return nil
}
