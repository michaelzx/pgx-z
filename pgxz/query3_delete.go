package pgxz

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

// Delete return RowsAffected and error
func Delete(db *PgDb, col ICol, options ...IOption) (int64, error) {
	if col.HasKey("delete_at") {
		col.Set("delete_at", time.Now())
		return Update(db, col, options...)
	}
	// real delete
	og := optionsToGroup(options)
	if len(og.wheres) == 0 {
		return 0, wrapErr(errors.New("og.wheres must > 0 "))
	}
	var sql strings.Builder
	sqlArgs := make([]any, 0)
	var sqlArgIdx int64
	var sqlArgAppend = func(v any) {
		sqlArgs = append(sqlArgs, v)
		sqlArgIdx++
	}
	// delete from
	sql.WriteString("delete from ")
	sql.WriteString("\"")
	sql.WriteString(col.TableName())
	sql.WriteString("\"")
	// where
	whereSql, whereArgs := resolveWheres(og.wheres...)
	for _, whereArg := range whereArgs {
		sqlArgAppend(whereArg)
		whereSql = strings.Replace(whereSql, "?", "$"+strconv.FormatInt(sqlArgIdx, 10), 1)
	}
	sql.WriteString(whereSql)
	sql.WriteString(";")
	// commit
	if DEBUG {
		debugPrint("delete", sql.String(), sqlArgs)
	}
	result, err := db.Exec(context.TODO(), sql.String(), sqlArgs...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
