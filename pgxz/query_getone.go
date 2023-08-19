package pgxz

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func GetOne[T IModel](db *PgDb, whereSql string, whereArgs ...any) (*T, error) {
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
	sql.WriteString(" where ")
	for _, whereArg := range whereArgs {
		sqlArgAppend(whereArg)
		whereSql = strings.Replace(whereSql, "?", "$"+strconv.FormatInt(sqlArgIdx, 10), 1)
	}
	sql.WriteString(whereSql)
	if !strings.Contains(whereSql, "limit") {
		sql.WriteString(" limit 1")
	}
	sql.WriteString(";")

	rows, _ := db.Query(context.TODO(), sql.String(), sqlArgs...)
	row, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByNameLax[T])
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, nil
	case err != nil:
		return nil, wrapErr(err)
	}
	return row, nil
}
