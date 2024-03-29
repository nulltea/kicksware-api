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
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/golang/glog"
	TLS "go.kicksware.com/api/shared/core/meta"
	"go.kicksware.com/api/shared/util"

	"go.kicksware.com/api/shared/core/meta"

	"go.kicksware.com/api/services/products/core/model"
	"go.kicksware.com/api/services/products/core/repo"
	"go.kicksware.com/api/services/products/usecase/business"
)

type repository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	timeout    time.Duration
}

func NewMongoRepository(config config.DataStoreConfig) (repo.SneakerProductRepository, error) {
	repo := &repository{
		timeout:  time.Duration(config.Timeout) * time.Second,
	}
	client, err := newMongoClient(config); if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepository")
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
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func newTLSConfig(tlsConfig *TLS.TLSCertificate) *tls.Config {
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

func (r *repository) Fetch(codes []string, params *meta.RequestParams) ([]*model.SneakerProduct, error) {
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

func (r *repository) FetchAll(params *meta.RequestParams) ([]*model.SneakerProduct, error) {
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

func (r *repository) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.SneakerProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := query.ToBson(); if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.FetchQuery")
	}
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

func (r *repository) Count(query meta.RequestQuery, params *meta.RequestParams) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := query.ToBson(); if err != nil {
		return 0, errors.Wrap(err, "repository.SneakerProduct.Count")
	}
	if err != nil {
		return 0, errors.Wrap(err, "repository.SneakerProduct.Count")
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, errors.Wrap(err, "repository.SneakerProduct.Count")
	}
	return int(count), nil
}

func (r *repository) CountAll() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter := bson.M{}
	count, err := r.collection.CountDocuments(ctx, filter); if err != nil {
		return 0, errors.Wrap(err, "repository.SneakerReference.Count")
	}
	return int(count), nil
}
