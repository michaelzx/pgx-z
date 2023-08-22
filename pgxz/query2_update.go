package pgxz

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

// Update return RowsAffected and error
func Update(db *PgDb, updates ICol, options ...IOption) (int64, error) {
	og := optionsToGroup(options)
	if len(og.wheres) == 0 {
		return 0, wrapErr(errors.New("og.wheres must > 0 "))
	}
	if !updates.IsSet("delete_at") {
		updates.Set("update_at", time.Now())
	}
	var sql strings.Builder
	sqlArgs := make([]any, 0)
	var sqlArgIdx int64
	var sqlArgAppend = func(v any) {
		sqlArgs = append(sqlArgs, v)
		sqlArgIdx++
	}
	// update what
	sql.WriteString("update ")
	sql.WriteString("\"")
	sql.WriteString(updates.TableName())
	sql.WriteString("\"")
	// set
	sql.WriteString(" set ")
	kIdx := 0
	for k, v := range updates.Mapping() {
		if kIdx != 0 {
			sql.WriteString(",")
		}
		sql.WriteString(k)
		sqlArgAppend(v)
		sql.WriteString("=$")
		sql.WriteString(strconv.FormatInt(sqlArgIdx, 10))
		kIdx++
	}
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
		debutPrint(sql.String(), sqlArgs)
	}
	result, err := db.Exec(context.TODO(), sql.String(), sqlArgs...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), err

}
