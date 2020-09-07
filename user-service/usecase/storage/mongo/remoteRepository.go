package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/model"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/repo"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/env"
)

type remoteRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	timeout    time.Duration
}

type remote struct {
	UserID   string             `bson:"user_id"`
	RemoteID string             `bson:"remote_id"`
	Provider model.UserProvider `bson:"provider"`
}

func NewRemoteRepository(config env.DataStoreConfig) (repo.RemoteRepository, error) {
	repo := &remoteRepository{
		timeout: time.Duration(config.Timeout) * time.Second,
	}
	client, err := newMongoClient(config
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRemoteRepository")
	}
	repo.client = client
	database := client.Database(config.Database)
	repo.database = database
	repo.collection = database.Collection(config.RemoteCollection)
	return repo, nil
}

func (r *remoteRepository) Connect(userID string, remoteID string, provider model.UserProvider) error {
	remote := &remote{
		userID,
		remoteID,
		provider,
	}
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, remote)
	if err != nil {
		return errors.Wrap(err, "repository.remote.Connect")
	}
	return nil
}

func (r *remoteRepository) Sync(userID string, remotes map[model.UserProvider]string) error {
	bulk := make([]interface{}, 0)

	for provider := range remotes {
		bulk = append(bulk, &remote{
			userID,
			remotes[provider],
			provider,
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	_, err := r.collection.InsertMany(ctx, bulk)
	if err != nil {
		return errors.Wrap(err, "repository.remote.Sync")
	}
	return nil
}

func (r *remoteRepository) Track(remoteID string, provider model.UserProvider) (string, error) {
	filter := bson.M{
		"remote_id": remoteID,
		"provider": provider,
	}

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	result := r.collection.FindOne(ctx, filter); if result == nil {
		return "", mongo.ErrNoDocuments
	}
	var remote remote
	if err := result.Decode(remote); err != nil {
		return "", err
	}
	return remote.UserID, nil
}


