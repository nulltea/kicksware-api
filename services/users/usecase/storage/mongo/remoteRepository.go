package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.kicksware.com/api/shared/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.kicksware.com/api/services/users/core/model"
	"go.kicksware.com/api/services/users/core/repo"
)

type remoteRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	timeout    time.Duration
}

func NewRemoteRepository(config config.DataStoreConfig) (repo.RemoteRepository, error) {
	repo := &remoteRepository{
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

func (r *remoteRepository) Connect(userID string, remoteID string, provider model.UserProvider) error {
	remote := &model.RemoteAuth{
		UserID:   userID,
		RemoteID: remoteID,
		Provider: provider,
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
		bulk = append(bulk, &model.RemoteAuth{
			UserID:   userID,
			RemoteID: remotes[provider],
			Provider: provider,
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
	var remote model.RemoteAuth
	if err := result.Decode(&remote); err != nil {
		return "", err
	}
	return remote.UserID, nil
}


