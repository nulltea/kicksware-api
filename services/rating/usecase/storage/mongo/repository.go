package mongo

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"time"

	"github.com/pkg/errors"
	"go.kicksware.com/api/shared/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/golang/glog"
	"go.kicksware.com/api/shared/core/meta"
	"go.kicksware.com/api/shared/util"

	"go.kicksware.com/api/services/rating/core/model"
	"go.kicksware.com/api/services/rating/core/repo"
	"go.kicksware.com/api/services/rating/usecase/business"
)

type repository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	timeout    time.Duration
}

func NewRepository(config config.DataStoreConfig) (repo.RatingRepository, error) {
	repo := &repository{
		timeout:  time.Duration(config.Timeout) * time.Second,
	}
	client, err := newMongoClient(config); if err != nil {
		return nil, errors.Wrap(err, "repository.NewRepository")
	}
	repo.client = client
	database := client.Database(config.Database)
	repo.database = database
	repo.collection = database.Collection(config.Collection)
	return repo, nil
}

func newMongoClient(config config.DataStoreConfig) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Timeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(config.URL),
	)
	err = client.Ping(ctx, readpref.Primary()); if err != nil {
		return nil, err
	}
	return client, nil
}

func newTLSConfig(tlsConfig *meta.TLSCertificate) *tls.Config {
	if !tlsConfig.EnableTLS {
		return nil
	}
	certs := x509.NewCertPool()
	pem, err := ioutil.ReadFile(tlsConfig.CertFile); if err != nil {
		glog.Fatalln(err)
	}
	certs.AppendCertsFromPEM(pem)
	return &tls.Config{
		RootCAs: certs,
	}
}

func (r *repository) FetchOne(code string, params *meta.RequestParams) (*model.Rating, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := r.buildQueryPipeline(bson.M{"entity_id": code}, params)
	cursor, err := r.collection.Aggregate(ctx, query); if err != nil {
		return nil, errors.Wrap(err, "repository.Rating.FetchOne")
	}
	defer cursor.Close(ctx)

	var rates []*model.Rating
	if err = cursor.All(ctx, &rates); err != nil {
		return nil, errors.Wrap(err, "repository.Rating.FetchOne")
	}
	if rates == nil || len(rates) == 0 {
		if err == mongo.ErrNoDocuments{
			return nil, errors.Wrap(business.ErrRatingNotFound, "repository.Rating.FetchOne")
		}
		return nil, errors.Wrap(err, "repository.Rating.FetchOne")
	}
	return rates[0], nil
}

func (r *repository) Fetch(codes []string, params *meta.RequestParams) ([]*model.Rating, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := r.buildQueryPipeline(bson.M{"entity_id": bson.M{"$in": codes}}, params)
	cursor, err := r.collection.Aggregate(ctx, query); if err != nil  {
		return nil, errors.Wrap(err, "repository.Rating.Fetch")
	}
	defer cursor.Close(ctx)

	var rates []*model.Rating
	if err = cursor.All(ctx, &rates); err != nil {
		return nil, errors.Wrap(err, "repository.Rating.Fetch")
	}
	if rates == nil || len(rates) == 0 {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrRatingNotFound, "repository.Rating.Fetch")
		}
		return nil, errors.Wrap(err, "repository.Rating.Fetch")
	}
	return rates, nil
}

func (r *repository) FetchAll(params *meta.RequestParams) ([]*model.Rating, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := r.buildQueryPipeline(bson.M{}, params)
	cursor, err := r.collection.Aggregate(ctx, query); if err != nil  {
		return nil, errors.Wrap(err, "repository.Rating.FetchAll")
	}
	defer cursor.Close(ctx)

	var rate []*model.Rating
	if err = cursor.All(ctx, &rate); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrRatingNotFound, "repository.Rating.FetchAll")
		}
		return nil, errors.Wrap(err, "repository.Rating.FetchAll")
	}
	return rate, nil
}

func (r *repository) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.Rating, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := query.ToBson(); if err != nil {
		return nil, errors.Wrap(err, "repository.Rating.FetchQuery")
	}
	queryPipe := r.buildQueryPipeline(filter, params)
	cursor, err := r.collection.Aggregate(ctx, queryPipe)
	if err != nil  {
		return nil, errors.Wrap(err, "repository.Rating.FetchQuery")
	}
	defer cursor.Close(ctx)

	var rates []*model.Rating
	if err = cursor.All(ctx, &rates); err != nil {
		return nil, errors.Wrap(err, "repository.Rating.FetchQuery")
	}
	if rates == nil || len(rates) == 0 {
		return nil, errors.Wrap(business.ErrRatingNotFound, "repository.Rating.FetchQuery")
	}
	return rates, nil
}

func (r *repository) StoreOne(rate *model.Rating) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, rate)
	if err != nil {
		return errors.Wrap(err, "repository.Rating.Store")
	}
	return nil
}

func (r *repository) Modify(rate *model.Rating) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	doc, err := util.ToBsonMap(rate); if err != nil {
		return errors.Wrap(err, "repository.Rating.Modify")
	}
	update := bson.D{
		{"$set", doc},
	}
	filter := bson.M{"unique_id": rate.UniqueID}
	if _, err = r.collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.Wrap(err, "repository.Rating.Modify")
	}
	return nil
}

func (r *repository) Remove(code string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"unique_id": code}
	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return errors.Wrap(err, "repository.Rating.Remove")
	}
	return nil
}

func (r *repository) Count(query meta.RequestQuery, params *meta.RequestParams) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := query.ToBson(); if err != nil {
		return 0, errors.Wrap(err, "repository.Rating.Count")
	}

	count, err := r.collection.CountDocuments(ctx, filter); if err != nil {
		return 0, errors.Wrap(err, "repository.Rating.Count")
	}
	return int(count), nil
}

func (r *repository) CountAll() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter := bson.M{}
	count, err := r.collection.CountDocuments(ctx, filter); if err != nil {
		return 0, errors.Wrap(err, "repository.Rating.Count")
	}
	return int(count), nil
}


func (r *repository) buildQueryPipeline(matchQuery bson.M, param *meta.RequestParams) mongo.Pipeline {
	pipe := mongo.Pipeline{}
	pipe = append(pipe, bson.D{{"$match", matchQuery}})

	return pipe
}

