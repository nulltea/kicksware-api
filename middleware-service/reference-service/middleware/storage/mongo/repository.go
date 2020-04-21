package mongo

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"reference-service/core/model"
	"reference-service/core/repo"
	"reference-service/middleware/business"
	"reference-service/util"
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
	sneakerReference := &model.SneakerReference{}
	filter := bson.M{"uniqueid": code}
	err := r.collection.FindOne(ctx, filter).Decode(&sneakerReference)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrReferenceNotFound, "repository.SneakerReference.FetchOne")
		}
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchOne")
	}
	return sneakerReference, nil
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

	var sneakerReferences []*model.SneakerReference
	if err = cursor.All(ctx, &sneakerReferences); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.Fetch")
	}
	if sneakerReferences == nil || len(sneakerReferences) == 0 {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrReferenceNotFound, "repository.SneakerReference.Fetch")
		}
		return nil, errors.Wrap(err, "repository.SneakerReference.Fetch")
	}
	return sneakerReferences, nil
}

func (r *repository) FetchAll() ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil  {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchAll")
	}
	defer cursor.Close(ctx)

	var sneakerReference []*model.SneakerReference
	if err = cursor.All(ctx, &sneakerReference); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrReferenceNotFound, "repository.SneakerReference.FetchAll")
		}
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchAll")
	}
	return sneakerReference, nil
}

func (r *repository) FetchQuery(query map[string]interface{}) ([]*model.SneakerReference, error) {
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

	var sneakerReferences []*model.SneakerReference
	if err = cursor.All(ctx, &sneakerReferences); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchQuery")
	}
	if sneakerReferences == nil || len(sneakerReferences) == 0 {
		return nil, errors.Wrap(business.ErrReferenceNotFound, "repository.SneakerReference.FetchQuery")
	}
	return sneakerReferences, nil
}

func (r *repository) StoreOne(sneakerReference *model.SneakerReference) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, sneakerReference)
	if err != nil {
		return errors.Wrap(err, "repository.SneakerReference.Store")
	}
	return nil
}

func (r *repository) Store(sneakerReferences []*model.SneakerReference) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	bulk := make([]interface{}, len(sneakerReferences))
	for i := range sneakerReferences {
		bulk[i] = sneakerReferences[i]
	}
	_, err := r.collection.InsertMany(ctx, bulk)
	if err != nil {
		return errors.Wrap(err, "repository.SneakerReference.Store")
	}
	return nil
}

func (r *repository) Modify(sneakerReference *model.SneakerReference) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	doc, err := util.ToBsonMap(sneakerReference)
	if err != nil {
		return errors.Wrap(err, "repository.SneakerReference.Modify")
	}
	update := bson.D{
		{"$set", doc},
	}
	filter := bson.M{"uniqueid": sneakerReference.UniqueId}
	if _, err = r.collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.Wrap(err, "repository.SneakerReference.Modify")
	}
	return nil
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