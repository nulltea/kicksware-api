package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/timoth-y/kicksware-platform/middleware-service/service-common/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/core/meta"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/core/model"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/core/repo"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/env"
	"github.com/timoth-y/kicksware-platform/middleware-service/order-service/usecase/business"
)

type repository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	timeout    time.Duration
}

func NewRepository(config env.DataStoreConfig) (repo.OrderRepository, error) {
	repo := &repository{
		timeout:  time.Duration(config.Timeout) * time.Second,
	}
	client, err := newMongoClient(config.URL, config.Timeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRepository")
	}
	repo.client = client
	database := client.Database(config.Database)
	repo.database = database
	repo.collection = database.Collection(config.Collection)
	return repo, nil
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL)); if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary()); if err != nil {
		return nil, err
	}
	return client, nil
}

func (r *repository) FetchOne(code string, params meta.RequestParams) (*model.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := r.buildQueryPipeline(bson.M{"unique_id": code}, params)
	cursor, err := r.collection.Aggregate(ctx, query); if err != nil {
		return nil, errors.Wrap(err, "repository.Order.FetchOne")
	}
	defer cursor.Close(ctx)

	var orders []*model.Order
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, errors.Wrap(err, "repository.Order.FetchOne")
	}
	if orders == nil || len(orders) == 0 {
		if err == mongo.ErrNoDocuments{
			return nil, errors.Wrap(business.ErrOrderNotFound, "repository.Order.FetchOne")
		}
		return nil, errors.Wrap(err, "repository.Order.FetchOne")
	}
	return orders[0], nil
}

func (r *repository) Fetch(codes []string, params meta.RequestParams) ([]*model.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := r.buildQueryPipeline(bson.M{"unique_id": bson.M{"$in": codes}}, params)
	cursor, err := r.collection.Aggregate(ctx, query); if err != nil  {
		return nil, errors.Wrap(err, "repository.Order.Fetch")
	}
	defer cursor.Close(ctx)

	var orders []*model.Order
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, errors.Wrap(err, "repository.Order.Fetch")
	}
	if orders == nil || len(orders) == 0 {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrOrderNotFound, "repository.Order.Fetch")
		}
		return nil, errors.Wrap(err, "repository.Order.Fetch")
	}
	return orders, nil
}

func (r *repository) FetchAll(params meta.RequestParams) ([]*model.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := r.buildQueryPipeline(bson.M{}, params)
	cursor, err := r.collection.Aggregate(ctx, query); if err != nil  {
		return nil, errors.Wrap(err, "repository.Order.FetchAll")
	}
	defer cursor.Close(ctx)

	var order []*model.Order
	if err = cursor.All(ctx, &order); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrOrderNotFound, "repository.Order.FetchAll")
		}
		return nil, errors.Wrap(err, "repository.Order.FetchAll")
	}
	return order, nil
}

func (r *repository) FetchQuery(query meta.RequestQuery, params meta.RequestParams) ([]*model.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := query.ToBson(); if err != nil {
		return nil, errors.Wrap(err, "repository.Order.FetchQuery")
	}
	queryPipe := r.buildQueryPipeline(filter, params)
	cursor, err := r.collection.Aggregate(ctx, queryPipe)
	if err != nil  {
		return nil, errors.Wrap(err, "repository.Order.FetchQuery")
	}
	defer cursor.Close(ctx)

	var orders []*model.Order
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, errors.Wrap(err, "repository.Order.FetchQuery")
	}
	if orders == nil || len(orders) == 0 {
		return nil, errors.Wrap(business.ErrOrderNotFound, "repository.Order.FetchQuery")
	}
	return orders, nil
}

func (r *repository) StoreOne(order *model.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, order)
	if err != nil {
		return errors.Wrap(err, "repository.Order.Store")
	}
	return nil
}

func (r *repository) Store(orders []*model.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	bulk := make([]interface{}, len(orders))
	for i := range orders {
		bulk[i] = orders[i]
	}
	_, err := r.collection.InsertMany(ctx, bulk)
	if err != nil {
		return errors.Wrap(err, "repository.Order.Store")
	}
	return nil
}

func (r *repository) Modify(order *model.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	doc, err := util.ToBsonMap(order); if err != nil {
		return errors.Wrap(err, "repository.Order.Modify")
	}
	update := bson.D{
		{"$set", doc},
	}
	filter := bson.M{"unique_id": order.UniqueID}
	if _, err = r.collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.Wrap(err, "repository.Order.Modify")
	}
	return nil
}

func (r *repository) Count(query meta.RequestQuery, params meta.RequestParams) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := query.ToBson(); if err != nil {
		return 0, errors.Wrap(err, "repository.Order.Count")
	}

	count, err := r.collection.CountDocuments(ctx, filter); if err != nil {
		return 0, errors.Wrap(err, "repository.Order.Count")
	}
	return int(count), nil
}

func (r *repository) CountAll() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter := bson.M{}
	count, err := r.collection.CountDocuments(ctx, filter); if err != nil {
		return 0, errors.Wrap(err, "repository.Order.Count")
	}
	return int(count), nil
}


func (r *repository) buildQueryPipeline(matchQuery bson.M, param meta.RequestParams) mongo.Pipeline {
	pipe := mongo.Pipeline{}
	pipe = append(pipe, bson.D{{"$match", matchQuery}})

	return pipe
}

