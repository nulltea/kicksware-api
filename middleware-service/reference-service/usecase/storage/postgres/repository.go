package postgres

import (
	"context"

	sqb "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/util"

	"reference-service/core/meta"
	"reference-service/core/model"
	"reference-service/core/repo"
	"reference-service/env"
	"reference-service/usecase/business"
)

type repository struct {
	db *sqlx.DB
	table string
}

func NewPostgresRepository(config env.DataStoreConfig) (repo.SneakerReferenceRepository, error) {
	db, err := newPostgresClient(config.URL)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewPostgresRepository")
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

func (r *repository) FetchOne(code string) (*model.SneakerReference, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sneakerReference := &model.SneakerReference{}
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(sqb.Eq{"UniqueId":code}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchOne")
	}
	if err = r.db.GetContext(ctx, sneakerReference, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchOne")
	}
	return sneakerReference, nil
}

func (r *repository) Fetch(codes []string, params meta.RequestParams) ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sneakerReferences := make([]*model.SneakerReference, 0)
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(sqb.Eq{"UniqueId":codes}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.Fetch")
	}
	if err = r.db.SelectContext(ctx, &sneakerReferences, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.Fetch")
	}
	if sneakerReferences == nil || len(sneakerReferences) == 0 {
		return nil, errors.Wrap(business.ErrReferenceNotFound, "repository.SneakerReference.Fetch")
	}
	return sneakerReferences, nil
}

func (r *repository) FetchAll(params meta.RequestParams) ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sneakerReferences := make([]*model.SneakerReference, 0)
	cmd, args, err := sqb.Select("*").From(r.table).PlaceholderFormat(sqb.Dollar).ToSql(); if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchAll")
	}
	if err = r.db.SelectContext(ctx, &sneakerReferences, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchAll")
	}
	if sneakerReferences == nil || len(sneakerReferences) == 0 {
		return nil, errors.Wrap(business.ErrReferenceNotFound, "repository.SneakerReference.FetchAll")
	}
	return sneakerReferences, nil
}

func (r *repository) FetchQuery(query meta.RequestQuery, params meta.RequestParams) ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sneakerReferences := make([]*model.SneakerReference, 0)
	where, err := query.ToSql(); if err != nil {
		return nil, err
	}
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(where).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchQuery")
	}
	if err = r.db.SelectContext(ctx, &sneakerReferences, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchQuery")
	}
	if sneakerReferences == nil || len(sneakerReferences) == 0 {
		return nil, errors.Wrap(business.ErrReferenceNotFound, "repository.SneakerReference.FetchQuery")
	}
	return sneakerReferences, nil
}

func (r *repository) StoreOne(sneakerReference *model.SneakerReference) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Insert(r.table).SetMap(util.ToMap(sneakerReference)).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.SneakerReference.Store")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.SneakerReference.Store")
	}
	return nil
}

func (r *repository) Store(sneakerReferences []*model.SneakerReference) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Insert(r.table).Values(sneakerReferences).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.SneakerReference.Store")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.SneakerReference.Store")
	}
	return nil
}

func (r *repository) Modify(sneakerReference *model.SneakerReference) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Update(r.table).SetMap(util.ToMap(sneakerReference)).
		Where(sqb.Eq{"UniqueId": sneakerReference.UniqueId}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.SneakerReference.Modify")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.SneakerReference.Modify")
	}
	return nil
}

func (r *repository) Count(query meta.RequestQuery, params meta.RequestParams) (count int, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	where, err := query.ToSql(); if err != nil {
		return 0, err
	}
	cmd, args, err := sqb.Select("COUNT(*)").From(r.table).
		Where(where).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "repository.SneakerReference.FetchQuery")
	}
	if err = r.db.SelectContext(ctx, &count, cmd, args); err != nil {
		return 0, errors.Wrap(err, "repository.SneakerReference.FetchQuery")
	}
	return
}

func (r *repository) CountAll() (count int, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Select("COUNT(*)").From(r.table).
		PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "repository.SneakerReference.FetchQuery")
	}
	if err = r.db.SelectContext(ctx, &count, cmd, args); err != nil {
		return 0, errors.Wrap(err, "repository.SneakerReference.FetchQuery")
	}
	return
}