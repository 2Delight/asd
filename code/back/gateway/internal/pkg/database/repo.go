package database

import (
	"context"
	"gateway-api/internal/pkg/database/schema"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"time"
)

var (
	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

type SpecRepository struct {
	pool *pgxpool.Pool
}

func NewSpecRepository(pool *pgxpool.Pool) *SpecRepository {
	return &SpecRepository{
		pool: pool,
	}
}

func (r *SpecRepository) GetSpecificationByID(ctx context.Context, id int64) (*schema.Specification, error) {
	query := psql.
		Select(
			"id",
			"name",
			"git_path",
			"status",
			"created_at",
			"updated_at",
		).From("specifications").
		Where(sq.Eq{"id": id}).
		Limit(1)

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "GetSpecificationByID.ToSql")
	}

	var specification schema.Specification
	err = pgxscan.Get(ctx, r.pool, &specification, rawSQL, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // Возвращаем nil, если запись не найдена
		}

		return nil, errors.Wrap(err, "GetSpecificationByID.Get")
	}

	return &specification, nil
}

func (r *SpecRepository) UpdateSpecificationStatus(ctx context.Context, id int64, newStatus string) error {
	query := psql.
		Update("specifications").
		Set("status", newStatus).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": id})

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "UpdateSpecificationStatus.ToSql")
	}

	_, err = r.pool.Exec(ctx, rawSQL, args...)
	if err != nil {
		return errors.Wrap(err, "UpdateSpecificationStatus.Exec")
	}

	return nil
}
