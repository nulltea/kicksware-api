package postgres

import (
	"context"

	sqb "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.kicksware.com/api/shared/config"
	"go.kicksware.com/api/shared/util"

	"go.kicksware.com/api/shared/core/meta"

	"go.kicksware.com/api/services/orders/core/model"
	"go.kicksware.com/api/services/orders/core/repo"
	"go.kicksware.com/api/services/orders/usecase/business"
)

type repository struct {
	db *sqlx.DB
	table string
}

func NewRepository(config config.DataStoreConfig) (repo.OrderRepository, error) {
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

func (r *repository) FetchOne(code string, params *meta.RequestParams) (*model.Order, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	order := &model.Order{}
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(sqb.Eq{"UniqueID":code}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Order.FetchOne")
	}
	if err = r.db.GetContext(ctx, order, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.Order.FetchOne")
	}
	return order, nil
}

func (r *repository) Fetch(codes []string, params *meta.RequestParams) ([]*model.Order, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	orders := make([]*model.Order, 0)
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(sqb.Eq{"UniqueID":codes}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Order.Fetch")
	}
	if err = r.db.SelectContext(ctx, &orders, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.Order.Fetch")
	}
	if orders == nil || len(orders) == 0 {
		return nil, errors.Wrap(business.ErrOrderNotFound, "repository.Order.Fetch")
	}
	return orders, nil
}

func (r *repository) FetchAll(params *meta.RequestParams) ([]*model.Order, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	orders := make([]*model.Order, 0)
	cmd, args, err := sqb.Select("*").From(r.table).PlaceholderFormat(sqb.Dollar).ToSql(); if err != nil {
		return nil, errors.Wrap(err, "repository.Order.FetchAll")
	}
	if err = r.db.SelectContext(ctx, &orders, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.Order.FetchAll")
	}
	if orders == nil || len(orders) == 0 {
		return nil, errors.Wrap(business.ErrOrderNotFound, "repository.Order.FetchAll")
	}
	return orders, nil
}

func (r *repository) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.Order, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	orders := make([]*model.Order, 0)
	where, err := query.ToSql(); if err != nil {
		return nil, err
	}
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(where).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Order.FetchQuery")
	}
	if err = r.db.SelectContext(ctx, &orders, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.Order.FetchQuery")
	}
	if orders == nil || len(orders) == 0 {
		return nil, errors.Wrap(business.ErrOrderNotFound, "repository.Order.FetchQuery")
	}
	return orders, nil
}

func (r *repository) StoreOne(order *model.Order) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Insert(r.table).SetMap(util.ToMap(order)).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.Order.Store")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.Order.Store")
	}
	return nil
}

func (r *repository) Modify(order *model.Order) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Update(r.table).SetMap(util.ToMap(order)).
		Where(sqb.Eq{"UniqueID": order.UniqueID}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.Order.Modify")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.Order.Modify")
	}
	return nil
}

func (r *repository) Remove(code string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Delete(r.table).Where(sqb.Eq{"UniqueId":code}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Remove")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Remove")
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
		return 0, errors.Wrap(err, "repository.Order.FetchQuery")
	}
	if err = r.db.SelectContext(ctx, &count, cmd, args); err != nil {
		return 0, errors.Wrap(err, "repository.Order.FetchQuery")
	}
	return
}

func (r *repository) CountAll() (count int, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Select("COUNT(*)").From(r.table).
		PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "repository.Order.FetchQuery")
	}
	if err = r.db.SelectContext(ctx, &count, cmd, args); err != nil {
		return 0, errors.Wrap(err, "repository.Order.FetchQuery")
	}
	return
}
