package postgres

import (
	"context"

	sqb "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.kicksware.com/api/service-common/config"
	"go.kicksware.com/api/service-common/util"

	"go.kicksware.com/api/service-common/core/meta"

	"go.kicksware.com/api/user-service/core/model"
	"go.kicksware.com/api/user-service/core/repo"
	"go.kicksware.com/api/user-service/usecase/business"
)

type repository struct {
	db *sqlx.DB
	table string
}

func NewRepository(config config.DataStoreConfig) (repo.UserRepository, error) {
	db, err := newPostgresClient(config.URL)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRepository")
	}
	repo := &repository{
		db: db,
		table:  config.Collection,
	}
	return repo, nil
}

func newPostgresClient(url string) (*sqlx.DB, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err := sqlx.ConnectContext(ctx,"pgx", url)
	if err != nil {
		return nil, errors.Wrap(err, "repository.newPostgresClient")
	}
	if err = db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "repository.newPostgresClient")
	}
	return db, nil
}

func (r *repository) FetchOne(code string) (*model.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	user := &model.User{}
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(sqb.Eq{"uniqueID":code}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchOne")
	}
	if err = r.db.GetContext(ctx, user, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchOne")
	}
	return user, nil
}

func (r *repository) Fetch(codes []string, params *meta.RequestParams) ([]*model.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	users := make([]*model.User, 0)
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(sqb.Eq{"uniqueID":codes}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.User.Fetch")
	}
	if err = r.db.SelectContext(ctx, &users, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.User.Fetch")
	}
	if users == nil || len(users) == 0 {
		return nil, errors.Wrap(business.ErrUserNotFound, "repository.User.Fetch")
	}
	return users, nil
}

func (r *repository) FetchAll(params *meta.RequestParams) ([]*model.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	users := make([]*model.User, 0)
	cmd, args, err := sqb.Select("*").From(r.table).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchAll")
	}
	if err = r.db.SelectContext(ctx, &users, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchAll")
	}
	if users == nil || len(users) == 0 {
		return nil, errors.Wrap(business.ErrUserNotFound, "repository.User.FetchAll")
	}
	return users, nil
}

func (r *repository) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	users := make([]*model.User, 0)
	where, err := query.ToSql(); if err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchQuery")
	}
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(where).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchQuery")
	}
	if err = r.db.SelectContext(ctx, &users, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchQuery")
	}
	if users == nil || len(users) == 0 {
		return nil, errors.Wrap(business.ErrUserNotFound, "repository.User.FetchQuery")
	}
	return users, nil
}

func (r *repository) Store(user *model.User) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Insert(r.table).SetMap(util.ToMap(user)).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.User.Store")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.User.Store")
	}
	return nil
}

func (r *repository) Modify(user *model.User) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Update(r.table).SetMap(util.ToMap(user)).
		Where(sqb.Eq{"uniqueID":user.UniqueID}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.User.Store")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.User.Store")
	}
	return nil
}

func (r *repository) Replace(user *model.User) error {
	if err := r.Modify(user); err != nil {
		return errors.Wrap(err, "repository.User.Replace")
	}
	return nil
}

func (r *repository) Remove(code string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Delete(r.table).Where(sqb.Eq{"UniqueID":code}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.User.Remove")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.User.Remove")
	}
	return nil
}

func (r *repository) Count(query meta.RequestQuery, params *meta.RequestParams) (count int, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	where, err := query.ToSql(); if err != nil {
		return 0, errors.Wrap(err, "repository.User.Count")
	}

	cmd, args, err := sqb.Select("COUNT(*)").From(r.table).
		Where(where).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "repository.User.Count")
	}
	if err = r.db.SelectContext(ctx, &count, cmd, args); err != nil {
		return 0, errors.Wrap(err, "repository.User.Count")
	}
	return
}

func (r *repository) CountAll() (count int, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd, args, err := sqb.Select("COUNT(*)").From(r.table).
		PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "repository.User.CountAll")
	}
	if err = r.db.SelectContext(ctx, &count, cmd, args); err != nil {
		return 0, errors.Wrap(err, "repository.User.CountAll")
	}
	return
}