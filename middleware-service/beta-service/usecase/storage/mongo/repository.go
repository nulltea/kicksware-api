package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"beta-service/core/meta"
	"beta-service/core/model"
	"beta-service/core/repo"
	"beta-service/env"
	"beta-service/usecase/business"
)

type repository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	timeout    time.Duration
}

func NewRepository(config env.DataStoreConfig) (repo.BetaRepository, error) {
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

func (r *repository) FetchOne(code string, params *meta.RequestParams) (*model.Beta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := r.buildQueryPipeline(bson.M{"unique_id": code}, params)
	cursor, err := r.collection.Aggregate(ctx, query); if err != nil {
		return nil, errors.Wrap(err, "repository.Beta.FetchOne")
	}
	defer cursor.Close(ctx)

	var betas []*model.Beta
	if err = cursor.All(ctx, &betas); err != nil {
		return nil, errors.Wrap(err, "repository.Beta.FetchOne")
	}
	if betas == nil || len(betas) == 0 {
		if err == mongo.ErrNoDocuments{
			return nil, errors.Wrap(business.ErrBetaNotFound, "repository.Beta.FetchOne")
		}
		return nil, errors.Wrap(err, "repository.Beta.FetchOne")
	}
	return betas[0], nil
}

func (r *repository) Fetch(codes []string, params *meta.RequestParams) ([]*model.Beta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := r.buildQueryPipeline(bson.M{"unique_id": bson.M{"$in": codes}}, params)
	cursor, err := r.collection.Aggregate(ctx, query); if err != nil  {
		return nil, errors.Wrap(err, "repository.Beta.Fetch")
	}
	defer cursor.Close(ctx)

	var betas []*model.Beta
	if err = cursor.All(ctx, &betas); err != nil {
		return nil, errors.Wrap(err, "repository.Beta.Fetch")
	}
	if betas == nil || len(betas) == 0 {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrBetaNotFound, "repository.Beta.Fetch")
		}
		return nil, errors.Wrap(err, "repository.Beta.Fetch")
	}
	return betas, nil
}

func (r *repository) FetchAll(params *meta.RequestParams) ([]*model.Beta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := r.buildQueryPipeline(bson.M{}, params)
	cursor, err := r.collection.Aggregate(ctx, query); if err != nil  {
		return nil, errors.Wrap(err, "repository.Beta.FetchAll")
	}
	defer cursor.Close(ctx)

	var beta []*model.Beta
	if err = cursor.All(ctx, &beta); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrBetaNotFound, "repository.Beta.FetchAll")
		}
		return nil, errors.Wrap(err, "repository.Beta.FetchAll")
	}
	return beta, nil
}

func (r *repository) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.Beta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := query.ToBson(); if err != nil {
		return nil, errors.Wrap(err, "repository.Beta.FetchQuery")
	}
	queryPipe := r.buildQueryPipeline(filter, params)
	cursor, err := r.collection.Aggregate(ctx, queryPipe)
	if err != nil  {
		return nil, errors.Wrap(err, "repository.Beta.FetchQuery")
	}
	defer cursor.Close(ctx)

	var betas []*model.Beta
	if err = cursor.All(ctx, &betas); err != nil {
		return nil, errors.Wrap(err, "repository.Beta.FetchQuery")
	}
	if betas == nil || len(betas) == 0 {
		return nil, errors.Wrap(business.ErrBetaNotFound, "repository.Beta.FetchQuery")
	}
	return betas, nil
}

func (r *repository) StoreOne(beta *model.Beta) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, beta)
	if err != nil {
		return errors.Wrap(err, "repository.Beta.Store")
	}
	return nil
}

func (r *repository) Store(betas []*model.Beta) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	bulk := make([]interface{}, len(betas))
	for i := range betas {
		bulk[i] = betas[i]
	}
	_, err := r.collection.InsertMany(ctx, bulk)
	if err != nil {
		return errors.Wrap(err, "repository.Beta.Store")
	}
	return nil
}

func (r *repository) Modify(beta *model.Beta) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	doc, err := util.ToBsonMap(beta); if err != nil {
		return errors.Wrap(err, "repository.Beta.Modify")
	}
	update := bson.D{
		{"$set", doc},
	}
	filter := bson.M{"unique_id": beta.UniqueID}
	if _, err = r.collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.Wrap(err, "repository.Beta.Modify")
	}
	return nil
}

func (r *repository) Count(query meta.RequestQuery, params *meta.RequestParams) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := query.ToBson(); if err != nil {
		return 0, errors.Wrap(err, "repository.Beta.Count")
	}

	count, err := r.collection.CountDocuments(ctx, filter); if err != nil {
		return 0, errors.Wrap(err, "repository.Beta.Count")
	}
	return int(count), nil
}

func (r *repository) CountAll() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter := bson.M{}
	count, err := r.collection.CountDocuments(ctx, filter); if err != nil {
		return 0, errors.Wrap(err, "repository.Beta.Count")
	}
	return int(count), nil
}


func (r *repository) buildQueryPipeline(matchQuery bson.M, param *meta.RequestParams) mongo.Pipeline {
	pipe := mongo.Pipeline{}
	pipe = append(pipe, bson.D{{"$match", matchQuery}})

	return pipe
}

