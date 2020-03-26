package mongo

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"elastic-search-service/core/model"
	"elastic-search-service/core/repo"
	"elastic-search-service/middleware/business"
	"elastic-search-service/util"
	"time"
)

type repository struct {
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

func NewMongoRepository(url, db, collection string, mongoTimeout int) (repo.SneakerReferenceRepository, error) {
	repo := &repository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
	}
	client, err := newMongoClient(url, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepo")
	}
	repo.client = client
	database := client.Database(db)
	repo.database = database
	repo.collection = database.Collection(collection)
	return repo, nil
}

func (r *repository) FetchOne(code string) (*model.SneakerReference, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	sneakerProduct := &model.SneakerReference{}
	filter := bson.M{"uniqueid": code}
	err := r.collection.FindOne(ctx, filter).Decode(&sneakerProduct)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrReferenceNotFound, "repository.SneakerReference.FetchOne")
		}
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchOne")
	}
	return sneakerProduct, nil
}

func (r *repository) Fetch(codes []string) ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"uniqueid": bson.M{"$in": codes}}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil  {
		return nil, errors.Wrap(err, "repository.SneakerReference.Fetch")
	}
	defer cursor.Close(ctx)

	var sneakerProducts []*model.SneakerReference
	if err = cursor.All(ctx, &sneakerProducts); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.Fetch")
	}
	if sneakerProducts == nil || len(sneakerProducts) == 0 {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrReferenceNotFound, "repository.SneakerReference.Fetch")
		}
		return nil, errors.Wrap(err, "repository.SneakerReference.Fetch")
	}
	return sneakerProducts, nil
}

func (r *repository) FetchAll() ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil  {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchAll")
	}
	defer cursor.Close(ctx)

	var sneakerProduct []*model.SneakerReference
	if err = cursor.All(ctx, &sneakerProduct); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrReferenceNotFound, "repository.SneakerReference.FetchAll")
		}
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchAll")
	}
	return sneakerProduct, nil
}

func (r *repository) FetchQuery(query interface{}) ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := util.ToBsonMap(query)
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchQuery")
	}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil  {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchQuery")
	}
	defer cursor.Close(ctx)

	var sneakerProducts []*model.SneakerReference
	if err = cursor.All(ctx, &sneakerProducts); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchQuery")
	}
	if sneakerProducts == nil || len(sneakerProducts) == 0 {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrReferenceNotFound, "repository.SneakerReference.FetchQuery")
		}
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchQuery")
	}
	return sneakerProducts, nil
}

func (r *repository) Count(query interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := util.ToBsonMap(query)
	if err != nil {
		return 0, errors.Wrap(err, "repository.SneakerReference.Count")
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, errors.Wrap(err, "repository.SneakerReference.Count")
	}
	return count, nil
}