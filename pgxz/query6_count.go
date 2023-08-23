package pgxz

import (
	"context"
	"strconv"
	"strings"
)

func Count[T IModel](db *PgDb, whereSql string, whereArgs ...any) (int64, error) {
	var t T
	var sql strings.Builder
	sqlArgs := make([]any, 0)
	var sqlArgIdx int64
	var sqlArgAppend = func(v any) {
		sqlArgs = append(sqlArgs, v)
		sqlArgIdx++
	}
	// select
	sql.WriteString("select count(*) from ")
	sql.WriteString("\"")
	sql.WriteString(t.TableName())
	sql.WriteString("\"")
	if whereSql != "" {
		// where
		sql.WriteString(" where ")
		for _, whereArg := range whereArgs {
			sqlArgAppend(whereArg)
			whereSql = strings.Replace(whereSql, "?", "$"+strconv.FormatInt(sqlArgIdx, 10), 1)
		}
		sql.WriteString(whereSql)
	}
	sql.WriteString(";")
	// commit
	if DEBUG {
		debugPrint("count", sql.String(), sqlArgs)
	}
	var c int64
	err := db.QueryRow(context.TODO(), sql.String(), sqlArgs...).Scan(&c)
	return c, err
}
