package mongo

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"middleware-serivce/model"
	"time"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
	collection string
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

func NewMongoRepository(mongoURL, mongoDB, collection string, mongoTimeout int) (model.SneakerProductRepository, error) {
	repo := &mongoRepository{
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

func (r *mongoRepository) Find(id string) (*model.SneakerProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	sneakerProduct := &model.SneakerProduct{}
	collection := r.client.Database(r.database).Collection(r.collection)
	filter := bson.M{"UniqueId": id}
	err := collection.FindOne(ctx, filter).Decode(&sneakerProduct)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(model.ErrProductNotFound, "repository.SneakerProduct.Find")
		}
		return nil, errors.Wrap(err, "repository.SneakerProduct.Find")
	}
	return sneakerProduct, nil
}

func (r *mongoRepository) Store(sneakerProduct *model.SneakerProduct) error {
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
