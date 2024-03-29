package postgres

import (
	"context"

	sqb "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.kicksware.com/api/shared/config"
	"go.kicksware.com/api/shared/util"

	"go.kicksware.com/api/services/users/core/model"
	"go.kicksware.com/api/services/users/core/repo"
)

type likesRepository struct {
	db *sqlx.DB
	table string
}

func NewLikesRepository(config config.DataStoreConfig) (repo.LikesRepository, error) {
	db, err := newPostgresClient(config.URL)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewLikesRepository")
	}
	repo := &likesRepository{
		db: db,
		table:  config.Collection,
	}
	return repo, nil
}

func (r *likesRepository) AddLike(userID string, entityID string) error {
	like := &model.Like{
		UserID:   userID,
		EntityID: entityID,
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Insert(r.table).SetMap(util.ToMap(like)).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.Likes.AddLike")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.Likes.AddLike")
	}
	return nil
}

func (r *likesRepository) RemoveLike(userID string, entityID string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Delete(r.table).Where(sqb.Eq{
		"UserID":userID,
		"EntityID": entityID,
	}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.Likes.RemoveLike")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.Likes.RemoveLike")
	}
	return nil
}
