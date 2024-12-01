package database

import (
	"context"
	"gateway-api/internal/pkg/database/schema"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/pkg/errors"
	"time"
)

// GetRecordsByIDs получает записи из указанной таблицы по списку ID.
// Возвращает слайс ссылок на model.Specification и ошибку, если она возникла.
func (r *SpecRepository) GetRecordsByIDs(ctx context.Context, tableName string, ids []int64) ([]*schema.Specification, error) {
	// 1. Построение SQL-запроса с использованием Squirrel
	query := psql.
		Select("id", "status").
		From(tableName).
		Where(sq.Eq{"id": ids})

	// 2. Преобразование запроса в SQL и получение аргументов
	rawSQL, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "GetRecordsByIDs.ToSql")
	}

	// 3. Выполнение запроса и сканирование результатов
	var specifications []*schema.Specification
	err = pgxscan.Select(ctx, r.pool, &specifications, rawSQL, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetRecordsByIDs.Select")
	}

	return specifications, nil
}

// UpdateRecordStatus обновляет статус записей в указанной таблице по списку ID.
// Возвращает ошибку, если операция не удалась.
func (r *SpecRepository) UpdateRecordStatus(ctx context.Context, tableName string, ids []int64, newStatus string) error {
	// 1. Построение UPDATE-запроса с использованием Squirrel
	query := psql.
		Update(tableName).
		Set("status", newStatus).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": ids})

	// 2. Преобразование запроса в SQL и получение аргументов
	rawSQL, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "UpdateRecordStatus.ToSql")
	}

	// 3. Выполнение запроса
	cmdTag, err := r.pool.Exec(ctx, rawSQL, args...)
	if err != nil {
		return errors.Wrap(err, "UpdateRecordStatus.Exec")
	}

	// 4. Проверка, что хотя бы одна запись была обновлена
	if cmdTag.RowsAffected() == 0 {
		return errors.New("записи не найдены или статус не изменён")
	}

	return nil
}
