package mongo

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"product-service/core/model"
	"product-service/core/repo"
	"product-service/scenario/business"
	"time"
)

type Repository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
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
	repo := &Repository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepo")
	}
	repo.client = client
	database := client.Database(mongoDB)
	repo.database = database
	repo.collection = database.Collection(collection)
	return repo, nil
}

func (r *Repository) RetrieveOne(code string) (*model.SneakerProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	sneakerProduct := &model.SneakerProduct{}
	filter := bson.M{"UniqueId": code}
	err := r.collection.FindOne(ctx, filter).Decode(&sneakerProduct)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.Retrieve")
		}
		return nil, errors.Wrap(err, "repository.SneakerProduct.Retrieve")
	}
	return sneakerProduct, nil
}

func (r *Repository) Retrieve(codes []string) ([]*model.SneakerProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"UniqueId": bson.M{"$in": codes}}

	cursor, err := r.collection.Find(ctx, filter)
	defer cursor.Close(ctx)

	var sneakerProduct []*model.SneakerProduct
	if err = cursor.All(ctx, &sneakerProduct); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.Retrieve")
	}
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.Retrieve")
		}
		return nil, errors.Wrap(err, "repository.SneakerProduct.Retrieve")
	}
	return sneakerProduct, nil
}

func (r *Repository) RetrieveAll() ([]*model.SneakerProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{}

	cursor, err := r.collection.Find(ctx, filter)
	defer cursor.Close(ctx)

	var sneakerProduct []*model.SneakerProduct
	if err = cursor.All(ctx, &sneakerProduct); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.Retrieve")
	}
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.Retrieve")
		}
		return nil, errors.Wrap(err, "repository.SneakerProduct.Retrieve")
	}
	return sneakerProduct, nil
}

func (r *Repository) RetrieveQuery(query interface{}) ([]*model.SneakerProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := query

	cursor, err := r.collection.Find(ctx, filter)
	defer cursor.Close(ctx)

	var sneakerProduct []*model.SneakerProduct
	if err = cursor.All(ctx, &sneakerProduct); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.Retrieve")
	}
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.Retrieve")
		}
		return nil, errors.Wrap(err, "repository.SneakerProduct.Retrieve")
	}
	return sneakerProduct, nil
}

func (r *Repository) Store(sneakerProduct *model.SneakerProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	_, err := r.collection.InsertOne(
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

func (r *Repository) Modify(sneakerProduct *model.SneakerProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	doc, err := bson.Marshal(sneakerProduct)
	if err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Store")
	}
	filter := bson.M{"UniqueId": sneakerProduct.UniqueId}
	if _, err = r.collection.UpdateOne(ctx, filter, doc); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Store")
	}
	return nil
}

func (r *Repository) Remove(sneakerProduct *model.SneakerProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"UniqueId": sneakerProduct.UniqueId}
	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Store")
	}
	return nil
}