package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.kicksware.com/api/service-common/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.kicksware.com/api/user-service/core/model"
	"go.kicksware.com/api/user-service/core/repo"
)

type subscriptionRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	timeout    time.Duration
}

func NewSubscriptionsRepository(config config.DataStoreConfig) (repo.SubscriptionRepository, error) {
	repo := &subscriptionRepository{
		timeout: time.Duration(config.Timeout) * time.Second,
	}
	client, err := newMongoClient(config); if err != nil {
		return nil, errors.Wrap(err, "repository.NewRemoteRepository")
	}
	repo.client = client
	database := client.Database(config.Database)
	repo.database = database
	repo.collection = database.Collection(config.Collection)
	return repo, nil
}

func (r *subscriptionRepository) Add(record model.MailSubscription) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, record)
	if err != nil {
		return errors.Wrap(err, "repository.subscription.Add")
	}
	return nil
}

func (r *subscriptionRepository) Delete(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"email": email}
	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return errors.Wrap(err, "repository.subscription.Delete")
	}
	return nil
}
