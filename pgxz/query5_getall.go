package pgxz

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func GetAll[T IModel](db *PgDb, col ICol, options ...IOption) ([]T, error) {
	og := optionsToGroup(options)
	// how to limit
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
		og.wheres = append(og.wheres, Where("delete_at is null"))
	}
	if len(og.wheres) > 0 {
		whereSql, whereArgs := resolveWheres(og.wheres...)
		for _, whereArg := range whereArgs {
			sqlArgAppend(whereArg)
			whereSql = strings.Replace(whereSql, "?", "$"+strconv.FormatInt(sqlArgIdx, 10), 1)
		}
		sql.WriteString(whereSql)
	}
	if og.limit != nil {
		sql.WriteString(og.limit.ToSql())
	}
	if og.offset != nil {
		sql.WriteString(og.offset.ToSql())
	}
	if og.group != nil {
		sql.WriteString(og.group.ToSql())
	}
	sql.WriteString(";")
	// commit
	if DEBUG {
		debugPrint("getall", sql.String(), sqlArgs)
	}
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
