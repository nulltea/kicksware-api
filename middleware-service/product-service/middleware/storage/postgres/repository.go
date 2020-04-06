package postgres

import (
	"context"

	sqb "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"product-service/core/model"
	"product-service/core/repo"
	"product-service/middleware/business"
	"product-service/util"
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

func NewPostgresRepository(connection, table string) (repo.SneakerProductRepository, error) {
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

func (r *repository) FetchOne(code string) (*model.SneakerProduct, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sneakerProduct := &model.SneakerProduct{}
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(sqb.Eq{"UniqueId":code}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.FetchOne")
	}
	if err = r.db.GetContext(ctx, sneakerProduct, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.FetchOne")
	}
	return sneakerProduct, nil
}

func (r *repository) Fetch(codes []string) ([]*model.SneakerProduct, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sneakerProducts := make([]*model.SneakerProduct, 0)
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(sqb.Eq{"UniqueId":codes}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.Fetch")
	}
	if err = r.db.SelectContext(ctx, &sneakerProducts, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.Fetch")
	}
	if sneakerProducts == nil || len(sneakerProducts) == 0 {
		return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.Fetch")
	}
	return sneakerProducts, nil
}

func (r *repository) FetchAll() ([]*model.SneakerProduct, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sneakerProducts := make([]*model.SneakerProduct, 0)
	cmd, args, err := sqb.Select("*").From(r.table).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.FetchAll")
	}
	if err = r.db.SelectContext(ctx, &sneakerProducts, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.FetchAll")
	}
	if sneakerProducts == nil || len(sneakerProducts) == 0 {
		return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.FetchAll")
	}
	return sneakerProducts, nil
}

func (r *repository) FetchQuery(query interface{}) ([]*model.SneakerProduct, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sneakerProducts := make([]*model.SneakerProduct, 0)
	where := util.ToSqlWhere(query)
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(where).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.FetchQuery")
	}
	if err = r.db.SelectContext(ctx, &sneakerProducts, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.SneakerProduct.FetchQuery")
	}
	if sneakerProducts == nil || len(sneakerProducts) == 0 {
		return nil, errors.Wrap(business.ErrProductNotFound, "repository.SneakerProduct.FetchQuery")
	}
	return sneakerProducts, nil
}

func (r *repository) Store(sneakerProduct *model.SneakerProduct) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Insert(r.table).SetMap(util.ToMap(sneakerProduct)).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Store")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Store")
	}
	return nil
}

func (r *repository) Modify(sneakerProduct *model.SneakerProduct) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Update(r.table).SetMap(util.ToMap(sneakerProduct)).
		Where(sqb.Eq{"UniqueId":sneakerProduct.UniqueId}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Store")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Store")
	}
	return nil
}

func (r *repository) Replace(sneakerProduct *model.SneakerProduct) error {
	if err := r.Modify(sneakerProduct); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.Replace")
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

func (r *repository) RemoveObj(sneakerProduct *model.SneakerProduct) error {
	if err := r.Remove(sneakerProduct.UniqueId); err != nil {
		return errors.Wrap(err, "repository.SneakerProduct.RemoveObj")
	}
	return nil
}

func (r *repository) Count(query interface{}) (count int64, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	where := util.ToSqlWhere(query)
	cmd, args, err := sqb.Select("COUNT(*)").From(r.table).
		Where(where).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "repository.SneakerProduct.FetchQuery")
	}
	if err = r.db.SelectContext(ctx, &count, cmd, args); err != nil {
		return 0, errors.Wrap(err, "repository.SneakerProduct.FetchQuery")
	}
	return
}