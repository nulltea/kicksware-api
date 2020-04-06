package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"product-service/core/model"
	"product-service/core/repo"
	"product-service/middleware/business"
	"product-service/util"
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

func NewMongoRepository(url, db, collection string, mongoTimeout int) (repo.SneakerProductRepository, error) {
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

func (r *repository) FetchOne(code string) (*model.SneakerProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	sneakerProduct := &model.SneakerProduct{}
	filter := bson.M{"uniqueid": code}
	err := r.collection.FindOne(ctx, filter).Decode(&sneakerProduct)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.FetchOne")
		}
		return nil, errors.Wrap(err, "repository.SneakerProduct.FetchOne")
	}
	return sneakerProduct, nil
}

func (r *repository) Fetch(codes []string) ([]*model.SneakerProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"uniqueid": bson.M{"$in": codes}}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil  {
		return nil, errors.Wrap(err, "repository.SneakerProduct.Fetch")
	}
	defer cursor.Close(ctx)

	var sneakerProducts []*model.SneakerProduct
	if err = cursor.All(ctx, &sneakerProducts); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.Fetch")
	}
	if sneakerProducts == nil || len(sneakerProducts) == 0 {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.Fetch")
		}
		return nil, errors.Wrap(err, "repository.SneakerProduct.Fetch")
	}
	return sneakerProducts, nil
}

func (r *repository) FetchAll() ([]*model.SneakerProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil  {
		return nil, errors.Wrap(err, "repository.SneakerProduct.FetchAll")
	}
	defer cursor.Close(ctx)

	var sneakerProduct []*model.SneakerProduct
	if err = cursor.All(ctx, &sneakerProduct); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.FetchAll")
		}
		return nil, errors.Wrap(err, "repository.SneakerProduct.FetchAll")
	}
	return sneakerProduct, nil
}

func (r *repository) FetchQuery(query interface{}) ([]*model.SneakerProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := util.ToBsonMap(query)
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.FetchQuery")
	}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil  {
		return nil, errors.Wrap(err, "repository.SneakerProduct.FetchQuery")
	}
	defer cursor.Close(ctx)

	var sneakerProducts []*model.SneakerProduct
	if err = cursor.All(ctx, &sneakerProducts); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.FetchQuery")
	}
	if sneakerProducts == nil || len(sneakerProducts) == 0 {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.FetchQuery")
		}
		return nil, errors.Wrap(err, "repository.SneakerProduct.FetchQuery")
	}
	return sneakerProducts, nil
}

func (r *repository) Store(sneakerProduct *model.SneakerProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, sneakerProduct)
	if err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Store")
	}
	return nil
}

func (r *repository) Modify(sneakerProduct *model.SneakerProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	doc, err := util.ToBsonMap(sneakerProduct)
	if err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Modify")
	}
	update := bson.D{
		{"$set", doc},
	}
	filter := bson.M{"uniqueid": sneakerProduct.UniqueId}
	if _, err = r.collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Modify")
	}
	return nil
}

func (r *repository) Replace(sneakerProduct *model.SneakerProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"uniqueid": sneakerProduct.UniqueId}
	if _, err := r.collection.ReplaceOne(ctx, filter, sneakerProduct); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Replace")
	}
	return nil
}

func (r *repository) Remove(code string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"uniqueid": code}
	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Remove")
	}
	return nil
}

func (r *repository) RemoveObj(sneakerProduct *model.SneakerProduct) error {
	if err := r.Remove(sneakerProduct.UniqueId); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.RemoveObj")
	}
	return nil
}

func (r *repository) Count(query interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := util.ToBsonMap(query)
	if err != nil {
		return 0, errors.Wrap(err, "repository.SneakerProduct.Count")
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, errors.Wrap(err, "repository.SneakerProduct.Count")
	}
	return count, nil
}