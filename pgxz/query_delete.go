package pgxz

import (
	"context"
	"strconv"
	"strings"
	"time"
)

func Delete(db *PgDb, col ICol, whereSql string, whereArgs ...any) error {
	if col.HasKey("delete_at") {
		col.Set("delete_at", time.Now())
		return Update(db, col, whereSql, whereArgs...)
	}
	var sql strings.Builder
	sqlArgs := make([]any, 0)
	var sqlArgIdx int64
	var sqlArgAppend = func(v any) {
		sqlArgs = append(sqlArgs, v)
		sqlArgIdx++
	}
	// update what
	sql.WriteString("delete from ")
	sql.WriteString("\"")
	sql.WriteString(col.TableName())
	sql.WriteString("\"")
	// where
	sql.WriteString(" where ")
	for _, whereArg := range whereArgs {
		sqlArgAppend(whereArg)
		whereSql = strings.Replace(whereSql, "?", "$"+strconv.FormatInt(sqlArgIdx, 10), 1)
	}
	sql.WriteString(whereSql)
	sql.WriteString(";")
	_, err := db.Exec(context.TODO(), sql.String(), sqlArgs...)
	return err
}
