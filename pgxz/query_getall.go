package pgxz

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func GetAll[T IModel](db *PgDb, col ICol, whereSql string, whereArgs ...any) ([]T, error) {
	var t T
	var sql strings.Builder
	sqlArgs := make([]any, 0)
	var sqlArgIdx int64
	var sqlArgAppend = func(v any) {
		sqlArgs = append(sqlArgs, v)
		sqlArgIdx++
	}
	// select
	sql.WriteString("select * from ")
	sql.WriteString("\"")
	sql.WriteString(t.TableName())
	sql.WriteString("\"")

	// where
	if strings.TrimSpace(whereSql) == "" {
		if col.HasKey("delete_at") {
			sql.WriteString(" where delete_at is null")
		}
	} else {
		sql.WriteString(" where ")
		if col.HasKey("delete_at") {
			whereSql = "delete_at is null and " + whereSql
		}
		for _, whereArg := range whereArgs {
			sqlArgAppend(whereArg)
			whereSql = strings.Replace(whereSql, "?", "$"+strconv.FormatInt(sqlArgIdx, 10), 1)
		}
		sql.WriteString(whereSql)
		if !strings.Contains(whereSql, "limit") {
			sql.WriteString(" limit 1")
		}
	}
	sql.WriteString(";")

	rows, _ := db.Query(context.TODO(), sql.String(), sqlArgs...)
	list, err := pgx.CollectRows[T](rows, pgx.RowToStructByNameLax[T])
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return []T{}, nil
	case err != nil:
		return nil, wrapErr(err)
	}
	return list, nil
}
