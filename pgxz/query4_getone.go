package pgxz

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func GetOne[T IModel](db *PgDb, col ICol, options ...IOption) (*T, error) {
	og := optionsToGroup(options)
	if len(og.wheres) == 0 {
		return nil, wrapErr(errors.New("wheres length must be > 0 "))
	}
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
	if col.HasKey("delete_at") {
		og.wheres = append(og.wheres, Where(" delete_at is null"))
	}
	whereSql, whereArgs := resolveWheres(og.wheres...)
	for _, whereArg := range whereArgs {
		sqlArgAppend(whereArg)
		whereSql = strings.Replace(whereSql, "?", "$"+strconv.FormatInt(sqlArgIdx, 10), 1)
	}
	sql.WriteString(whereSql)
	if !strings.Contains(whereSql, "limit") {
		sql.WriteString(" limit 1")
	}
	sql.WriteString(";")
	// commit
	debutPrint(&sql, sqlArgs)
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
