package pgxz

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

func Update(db *PgDb, updates ICol, whereSql string, whereArgs ...any) error {
	whereSql = strings.TrimSpace(whereSql)
	if whereSql == "" || len(whereArgs) == 0 {
		return errors.New("whereSql or whereArgs must have value")
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
