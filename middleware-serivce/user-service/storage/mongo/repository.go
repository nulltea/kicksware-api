package mongo

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
	"user-service/core/model"
	"user-service/core/repo"
	"user-service/scenario/business"
)

type MongoRepository struct {
	client     *mongo.Client
	database   string
	collection string
	timeout    time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoRepository(mongoURL, mongoDB, collection string, mongoTimeout int) (repo.SneakerProductRepository, error) {
	repo := &MongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
		collection: collection,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepo")
	}
	repo.client = client
	return repo, nil
}

func (r *MongoRepository) Retrieve(id string) (*model.SneakerProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	sneakerProduct := &model.SneakerProduct{}
	collection := r.client.Database(r.database).Collection(r.collection)
	filter := bson.M{"UniqueId": id}
	err := collection.FindOne(ctx, filter).Decode(&sneakerProduct)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.Retrieve")
		}
		return nil, errors.Wrap(err, "repository.SneakerProduct.Retrieve")
	}
	return sneakerProduct, nil
}

func (r *MongoRepository) Store(sneakerProduct *model.SneakerProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"UniqueId":        sneakerProduct.UniqueId,
			"ModelName": sneakerProduct.BrandName,
			"BrandName": sneakerProduct.ModelName,
			"Price": sneakerProduct.Price,
			"Type": sneakerProduct.Type,
			"Color": sneakerProduct.Color,
			"Condition": sneakerProduct.Condition,
			"Description": sneakerProduct.Description,
			"Owner":     sneakerProduct.Owner,
			"ConditionIndex": sneakerProduct.ConditionIndex,
			"AddedAt":   sneakerProduct.AddedAt,
		},
	)
	if err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Store")
	}
	return nil
}
