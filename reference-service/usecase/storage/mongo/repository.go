package mongo

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"io/ioutil"
	"time"

	"github.com/golang/glog"
	TLS "github.com/timoth-y/kicksware-api/service-common/core/meta"
	"github.com/timoth-y/kicksware-api/service-common/util"

	"github.com/timoth-y/kicksware-api/reference-service/core/meta"
	"github.com/timoth-y/kicksware-api/reference-service/core/model"
	"github.com/timoth-y/kicksware-api/reference-service/core/repo"
	"github.com/timoth-y/kicksware-api/reference-service/env"
	"github.com/timoth-y/kicksware-api/reference-service/usecase/business"
)

type repository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	timeout    time.Duration
}

func NewMongoRepository(config env.DataStoreConfig) (repo.SneakerReferenceRepository, error) {
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

func newMongoClient(config env.DataStoreConfig) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Timeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(config.URL).
		SetTLSConfig(newTLSConfig(config.TLS)).
		SetAuth(options.Credential{
			Username: config.Login, Password: config.Password,
		}),
	)
	err = client.Ping(ctx, readpref.Primary()); if err != nil {
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

func (r *repository) FetchOne(code string, params *meta.RequestParams) (*model.SneakerReference, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := r.buildQueryPipeline(bson.M{"uniqueid": code}, params)
	cursor, err := r.collection.Aggregate(ctx, query); if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchOne")
	}
	defer cursor.Close(ctx)

	var sneakerReferences []*model.SneakerReference
	if err = cursor.All(ctx, &sneakerReferences); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchOne")
	}
	if sneakerReferences == nil || len(sneakerReferences) == 0 {
		if err == mongo.ErrNoDocuments{
			return nil, errors.Wrap(business.ErrReferenceNotFound, "repository.SneakerReference.FetchOne")
		}
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchOne")
	}
	return sneakerReferences[0], nil
}

func (r *repository) Fetch(codes []string, params *meta.RequestParams) ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := r.buildQueryPipeline(bson.M{"uniqueid": bson.M{"$in": codes}}, params)
	cursor, err := r.collection.Aggregate(ctx, query); if err != nil  {
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

func (r *repository) FetchAll(params *meta.RequestParams) ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := r.buildQueryPipeline(bson.M{}, params)
	cursor, err := r.collection.Aggregate(ctx, query); if err != nil  {
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

func (r *repository) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := query.ToBson(); if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchQuery")
	}
	queryPipe := r.buildQueryPipeline(filter, params)
	cursor, err := r.collection.Aggregate(ctx, queryPipe)
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
	doc, err := util.ToBsonMap(sneakerReference); if err != nil {
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

func (r *repository) Count(query meta.RequestQuery, params *meta.RequestParams) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := query.ToBson(); if err != nil {
		return 0, errors.Wrap(err, "repository.SneakerReference.Count")
	}

	count, err := r.collection.CountDocuments(ctx, filter); if err != nil {
		return 0, errors.Wrap(err, "repository.SneakerReference.Count")
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


func (r *repository) buildQueryPipeline(matchQuery bson.M, param *meta.RequestParams) mongo.Pipeline {
	pipe := mongo.Pipeline{}
	pipe = append(pipe, bson.D{{"$match", matchQuery}})

	pipe = append(pipe, bson.D{
		{"$lookup", bson.M {
			"from": "likes",
			"localField": "uniqueid",
			"foreignField": "entity_id",
			"as": "like",
		}},
	})

	pipe = append(pipe, bson.D {
		{ "$addFields", bson.M {
			"likes": bson.M{"$size": "$like"},
		}},
	})

	if param != nil {
		if param.SortBy() != "" {
			pipe = append(pipe, bson.D {
				{"$sort", bson.M {param.SortBy(): param.SortDirectionNum(), "added_date": -1 }},
			})
		}
		if param.Offset() != 0 {
			pipe = append(pipe, bson.D {
				{"$skip", param.Offset()},
			})
		}
		if param.Limit() != 0 {
			pipe = append(pipe, bson.D {
				{"$limit",  param.Limit()},
			})
		}
	}

	if param != nil && len(param.UserID()) != 0 {
		pipe = append(pipe, bson.D {
			{ "$addFields", bson.M {
				"liked": bson.M{"$in": bson.A{param.UserID(),"$like.user_id"}},
			}},
		})
	}

	pipe = append(pipe, bson.D {
		{ "$project", bson.M {
			"like": 0,
		}},
	})

	pipe = append(pipe, bson.D {
		{"$lookup", bson.M {
			"from": "brands",
			"localField": "brand",
			"foreignField": "uniqueid",
			"as": "brand",
		}},
	})
	pipe = append(pipe, bson.D {{ "$unwind", bson.M{
		"path": "$brand",
		"preserveNullAndEmptyArrays": true,
	}},
	})

	pipe = append(pipe, bson.D {
		{"$lookup", bson.M {
			"from": "models",
			"localField": "model",
			"foreignField": "uniqueid",
			"as": "model",
		}},
	})
	pipe = append(pipe, bson.D {{ "$unwind", bson.M{
		"path": "$model",
		"preserveNullAndEmptyArrays": true,
	}},
	})

	pipe = append(pipe, bson.D{
		{"$lookup", bson.M {
			"from": "models",
			"localField": "basemodel",
			"foreignField": "uniqueid",
			"as": "basemodel",
		}},
	})
	pipe = append(pipe, bson.D{{ "$unwind", bson.M{
		"path": "$basemodel",
		"preserveNullAndEmptyArrays": true,
	}},
	})

	return pipe
}

