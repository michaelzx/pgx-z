package pgxz

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func Raw[T any](db *PgDb, sqlStr string, sqlArgs ...any) ([]T, error) {
	rows, err := db.Query(context.TODO(), sqlStr, sqlArgs...)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return []T{}, nil
	case err != nil:
		return nil, err
	}

	list := make([]T, 0)
	for rows.Next() {
		var v T
		if err := rows.Scan(&v); err != nil {
			return nil, err
		}
		list = append(list, v)
	}
	return list, err
}
