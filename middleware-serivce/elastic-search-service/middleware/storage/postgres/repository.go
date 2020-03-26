package postgres

import (
	"context"
	sqb "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"elastic-search-service/core/model"
	"elastic-search-service/core/repo"
	"elastic-search-service/middleware/business"
	"elastic-search-service/util"
)

type repository struct {
	db *sqlx.DB
	table string
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

func NewPostgresRepository(connection, table string) (repo.SneakerReferenceRepository, error) {
	db, err := newPostgresClient(connection)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewPostgresRepo")
	}
	repo := &repository{
		db: db,
		table:  table,
	}
	return repo, nil
}

func (r *repository) FetchOne(code string) (*model.SneakerReference, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sneakerProduct := &model.SneakerReference{}
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(sqb.Eq{"UniqueId":code}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchOne")
	}
	if err = r.db.GetContext(ctx, sneakerProduct, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchOne")
	}
	return sneakerProduct, nil
}

func (r *repository) Fetch(codes []string) ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sneakerProducts := make([]*model.SneakerReference, 0)
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(sqb.Eq{"UniqueId":codes}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.Fetch")
	}
	if err = r.db.SelectContext(ctx, &sneakerProducts, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.Fetch")
	}
	if sneakerProducts == nil || len(sneakerProducts) == 0 {
		return nil, errors.Wrap(business.ErrReferenceNotFound, "repository.SneakerReference.Fetch")
	}
	return sneakerProducts, nil
}

func (r *repository) FetchAll() ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sneakerProducts := make([]*model.SneakerReference, 0)
	cmd, args, err := sqb.Select("*").From(r.table).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchAll")
	}
	if err = r.db.SelectContext(ctx, &sneakerProducts, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchAll")
	}
	if sneakerProducts == nil || len(sneakerProducts) == 0 {
		return nil, errors.Wrap(business.ErrReferenceNotFound, "repository.SneakerReference.FetchAll")
	}
	return sneakerProducts, nil
}

func (r *repository) FetchQuery(query interface{}) ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sneakerProducts := make([]*model.SneakerReference, 0)
	where := util.ToSqlWhere(query)
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(where).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchQuery")
	}
	if err = r.db.SelectContext(ctx, &sneakerProducts, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerReference.FetchQuery")
	}
	if sneakerProducts == nil || len(sneakerProducts) == 0 {
		return nil, errors.Wrap(business.ErrReferenceNotFound, "repository.SneakerReference.FetchQuery")
	}
	return sneakerProducts, nil
}

func (r *repository) Count(query interface{}) (count int64, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	where := util.ToSqlWhere(query)
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