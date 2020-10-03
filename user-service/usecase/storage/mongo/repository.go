package mongo

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"time"

	"github.com/pkg/errors"
	"go.kicksware.com/api/service-common/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/golang/glog"
	"go.kicksware.com/api/service-common/core/meta"
	"go.kicksware.com/api/service-common/util"

	"go.kicksware.com/api/user-service/core/model"
	"go.kicksware.com/api/user-service/core/repo"
	"go.kicksware.com/api/user-service/usecase/business"
)

type repository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	remoteCollection *mongo.Collection
	timeout    time.Duration
}

func NewRepository(config config.DataStoreConfig) (repo.UserRepository, error) {
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
		ApplyURI(config.URL).
		SetTLSConfig(newTLSConfig(config.TLS)).
		SetAuth(options.Credential{
			Username: config.Login, Password: config.Password,
		}),
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

func (r *repository) FetchOne(userID string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	query := r.buildQueryPipeline(bson.M{"unique_id": userID}, nil)
	cursor, err := r.collection.Aggregate(ctx, query); if err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchOne")
	}
	defer cursor.Close(ctx)

	var users []*model.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchOne")
	}
	if err != nil || len(users) == 0 {
		if err == mongo.ErrNoDocuments || len(users) == 0 {
			return nil, errors.Wrap(business.ErrUserNotFound, "repository.User.FetchOne")
		}
		return nil, errors.Wrap(err, "repository.User.FetchOne")
	}
	return users[0], nil
}

func (r *repository) Fetch(usernames []string, params *meta.RequestParams) ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter := bson.M{"unique_id": bson.M{"$in": usernames}}
	query := r.buildQueryPipeline(filter, nil)
	cursor, err := r.collection.Aggregate(ctx, query); if err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchOne")
	}
	defer cursor.Close(ctx)

	var users []*model.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchOne")
	}
	if err = cursor.All(ctx, &users); err != nil {
		return nil, errors.Wrap(err, "repository.User.Fetch")
	}
	if users == nil || len(users) == 0 {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrUserNotFound, "repository.User.Fetch")
		}
		return nil, errors.Wrap(err, "repository.User.Fetch")
	}
	return users, nil
}

func (r *repository) FetchAll(params *meta.RequestParams) ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := r.buildQueryPipeline(bson.M{}, nil)
	cursor, err := r.collection.Aggregate(ctx, query); if err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchOne")
	}
	defer cursor.Close(ctx)

	var user []*model.User
	if err = cursor.All(ctx, &user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrUserNotFound, "repository.User.FetchAll")
		}
		return nil, errors.Wrap(err, "repository.User.FetchAll")
	}
	return user, nil
}

func (r *repository) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := query.ToBson()
	if err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchQuery")
	}
	queryPipe := r.buildQueryPipeline(filter, nil)
	cursor, err := r.collection.Aggregate(ctx, queryPipe); if err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchOne")
	}

	var users []*model.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchQuery")
	}
	if users == nil || len(users) == 0 {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrUserNotFound, "repository.User.FetchQuery")
		}
		return nil, errors.Wrap(err, "repository.User.FetchQuery")
	}
	return users, nil
}

func (r *repository) Store(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return errors.Wrap(err, "repository.User.Store")
	}
	return nil
}

func (r *repository) Modify(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	doc, err := util.ToBsonMap(user)
	if err != nil {
		return errors.Wrap(err, "repository.User.Modify")
	}
	update := bson.D{
		{"$set", doc},
	}
	filter := bson.M{"unique_id": user.UniqueID}
	if _, err = r.collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.Wrap(err, "repository.User.Modify")
	}
	return nil
}

func (r *repository) Replace(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"unique_id": user.UniqueID}
	if _, err := r.collection.ReplaceOne(ctx, filter, user); err != nil {
		return errors.Wrap(err, "repository.User.Replace")
	}
	return nil
}

func (r *repository) Remove(code string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"unique_id": code}
	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return errors.Wrap(err, "repository.User.Remove")
	}
	return nil
}

func (r *repository) Count(query meta.RequestQuery, params *meta.RequestParams) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := query.ToBson(); if err != nil {
		return 0, errors.Wrap(err, "repository.User.Count")
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, errors.Wrap(err, "repository.User.Count")
	}
	return int(count), nil
}

func (r *repository) CountAll() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter := bson.M{}
	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, errors.Wrap(err, "repository.User.Count")
	}
	return int(count), nil
}

func (r *repository) buildQueryPipeline(matchQuery bson.M, param *meta.RequestParams) mongo.Pipeline {
	pipe := mongo.Pipeline{}
	pipe = append(pipe, bson.D{{"$match", matchQuery}})

	pipe = append(pipe, bson.D {
		{"$lookup", bson.M {
			"from": "likes",
			"localField": "unique_id",
			"foreignField": "user_id",
			"as": "like",
		}},
	})
	pipe = append(pipe, bson.D {
		{ "$addFields", bson.M {
			"liked": "$like.entity_id",
		}},
	})

	pipe = append(pipe, bson.D {
		{ "$project", bson.M {
				"like": 0,
		}},
	})

	return pipe
}
