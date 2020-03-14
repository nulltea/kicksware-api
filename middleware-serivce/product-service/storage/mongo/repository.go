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
	"product-service/tool"
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
			return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.RetrieveOne")
		}
		return nil, errors.Wrap(err, "repository.SneakerProduct.RetrieveOne")
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

	cursor, err := r.collection.Find(ctx, bson.D{})
	defer cursor.Close(ctx)

	var sneakerProduct []*model.SneakerProduct
	if err = cursor.All(ctx, &sneakerProduct); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.RetrieveAll")
		}
		return nil, errors.Wrap(err, "repository.SneakerProduct.RetrieveAll")
	}
	return sneakerProduct, nil
}

func (r *Repository) RetrieveQuery(query interface{}) ([]*model.SneakerProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := tool.ToBsonMap(query)
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.RetrieveQuery")
	}
	cursor, err := r.collection.Find(ctx, filter)
	defer cursor.Close(ctx)

	var sneakerProduct []*model.SneakerProduct
	if err = cursor.All(ctx, &sneakerProduct); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.RetrieveQuery")
	}
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.RetrieveQuery")
		}
		return nil, errors.Wrap(err, "repository.SneakerProduct.RetrieveQuery")
	}
	return sneakerProduct, nil
}

func (r *Repository) Store(sneakerProduct *model.SneakerProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, sneakerProduct)
	if err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Store")
	}
	return nil
}

func (r *Repository) Modify(sneakerProduct *model.SneakerProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	doc, err := tool.ToBsonMap(sneakerProduct)
	if err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Modify")
	}
	update := bson.D{
		{"$set", doc},
	}
	filter := bson.M{"UniqueId": sneakerProduct.UniqueId}
	if _, err = r.collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Modify")
	}
	return nil
}

func (r *Repository) Replace(sneakerProduct *model.SneakerProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"UniqueId": sneakerProduct.UniqueId}
	if _, err := r.collection.ReplaceOne(ctx, filter, sneakerProduct); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Replace")
	}
	return nil
}

func (r *Repository) Remove(code string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"UniqueId": code}
	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Remove")
	}
	return nil
}

func (r *Repository) RemoveObj(sneakerProduct *model.SneakerProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"UniqueId": sneakerProduct.UniqueId}
	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.RemoveObj")
	}
	return nil
}

func (r *Repository) Count(query interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := tool.ToBsonMap(query)
	if err != nil {
		return 0, errors.Wrap(err, "repository.SneakerProduct.Count")
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, errors.Wrap(err, "repository.SneakerProduct.Count")
	}
	return count, nil
}