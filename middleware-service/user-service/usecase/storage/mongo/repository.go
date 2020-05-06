package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"user-service/core/meta"
	"user-service/core/model"
	"user-service/core/repo"
	"user-service/usecase/business"
	"user-service/util"
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

func NewMongoRepository(url, db, collection string, mongoTimeout int) (repo.UserRepository, error) {
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

func (r *repository) FetchOne(username string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	user := &model.User{}
	filter := bson.M{"username": username}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(business.ErrUserNotFound, "repository.User.FetchOne")
		}
		return nil, errors.Wrap(err, "repository.User.FetchOne")
	}
	return user, nil
}

func (r *repository) Fetch(usernames []string) ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"username": bson.M{"$in": usernames}}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil  {
		return nil, errors.Wrap(err, "repository.User.Fetch")
	}
	defer cursor.Close(ctx)

	var users []*model.User
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

func (r *repository) FetchAll() ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil  {
		return nil, errors.Wrap(err, "repository.User.FetchAll")
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

func (r *repository) FetchQuery(query meta.RequestQuery) ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter, err := query.ToBson()
	if err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchQuery")
	}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil  {
		return nil, errors.Wrap(err, "repository.User.FetchQuery")
	}
	defer cursor.Close(ctx)

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
	filter := bson.M{"UniqueId": user.UniqueId}
	if _, err = r.collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.Wrap(err, "repository.User.Modify")
	}
	return nil
}

func (r *repository) Replace(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"uniqueid": user.UniqueId}
	if _, err := r.collection.ReplaceOne(ctx, filter, user); err != nil {
		return errors.Wrap(err, "repository.User.Replace")
	}
	return nil
}

func (r *repository) Remove(code string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"UniqueId": code}
	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return errors.Wrap(err, "repository.User.Remove")
	}
	return nil
}

func (r *repository) Count(query meta.RequestQuery) (int, error) {
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