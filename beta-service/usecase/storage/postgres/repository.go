package postgres

import (
	"context"

	sqb "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/timoth-y/kicksware-api/service-common/util"

	"github.com/timoth-y/kicksware-api/beta-service/core/meta"
	"github.com/timoth-y/kicksware-api/beta-service/core/model"
	"github.com/timoth-y/kicksware-api/beta-service/core/repo"
	"github.com/timoth-y/kicksware-api/beta-service/env"
	"github.com/timoth-y/kicksware-api/beta-service/usecase/business"
)

type repository struct {
	db *sqlx.DB
	table string
}

func NewRepository(config env.DataStoreConfig) (repo.BetaRepository, error) {
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

func (r *repository) FetchOne(code string, params *meta.RequestParams) (*model.Beta, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	beta := &model.Beta{}
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(sqb.Eq{"UniqueID":code}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Beta.FetchOne")
	}
	if err = r.db.GetContext(ctx, beta, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.Beta.FetchOne")
	}
	return beta, nil
}

func (r *repository) Fetch(codes []string, params *meta.RequestParams) ([]*model.Beta, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	betas := make([]*model.Beta, 0)
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(sqb.Eq{"UniqueID":codes}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Beta.Fetch")
	}
	if err = r.db.SelectContext(ctx, &betas, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.Beta.Fetch")
	}
	if betas == nil || len(betas) == 0 {
		return nil, errors.Wrap(business.ErrBetaNotFound, "repository.Beta.Fetch")
	}
	return betas, nil
}

func (r *repository) FetchAll(params *meta.RequestParams) ([]*model.Beta, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	betas := make([]*model.Beta, 0)
	cmd, args, err := sqb.Select("*").From(r.table).PlaceholderFormat(sqb.Dollar).ToSql(); if err != nil {
		return nil, errors.Wrap(err, "repository.Beta.FetchAll")
	}
	if err = r.db.SelectContext(ctx, &betas, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.Beta.FetchAll")
	}
	if betas == nil || len(betas) == 0 {
		return nil, errors.Wrap(business.ErrBetaNotFound, "repository.Beta.FetchAll")
	}
	return betas, nil
}

func (r *repository) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.Beta, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	betas := make([]*model.Beta, 0)
	where, err := query.ToSql(); if err != nil {
		return nil, err
	}
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(where).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Beta.FetchQuery")
	}
	if err = r.db.SelectContext(ctx, &betas, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.Beta.FetchQuery")
	}
	if betas == nil || len(betas) == 0 {
		return nil, errors.Wrap(business.ErrBetaNotFound, "repository.Beta.FetchQuery")
	}
	return betas, nil
}

func (r *repository) StoreOne(beta *model.Beta) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Insert(r.table).SetMap(util.ToMap(beta)).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.Beta.Store")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.Beta.Store")
	}
	return nil
}

func (r *repository) Store(betas []*model.Beta) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Insert(r.table).Values(betas).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.Beta.Store")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.Beta.Store")
	}
	return nil
}

func (r *repository) Modify(beta *model.Beta) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Update(r.table).SetMap(util.ToMap(beta)).
		Where(sqb.Eq{"UniqueID": beta.UniqueID}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.Beta.Modify")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.Beta.Modify")
	}
	return nil
}

func (r *repository) Count(query meta.RequestQuery, params *meta.RequestParams) (count int, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	where, err := query.ToSql(); if err != nil {
		return 0, err
	}
	cmd, args, err := sqb.Select("COUNT(*)").From(r.table).
		Where(where).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "repository.Beta.FetchQuery")
	}
	if err = r.db.SelectContext(ctx, &count, cmd, args); err != nil {
		return 0, errors.Wrap(err, "repository.Beta.FetchQuery")
	}
	return
}

func (r *repository) CountAll() (count int, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Select("COUNT(*)").From(r.table).
		PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "repository.Beta.FetchQuery")
	}
	if err = r.db.SelectContext(ctx, &count, cmd, args); err != nil {
		return 0, errors.Wrap(err, "repository.Beta.FetchQuery")
	}
	return
}