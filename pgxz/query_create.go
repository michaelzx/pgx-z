package pgxz

import (
	"context"
	"strconv"
	"strings"
)

func Create(db *PgDb, col ICol) error {
	var sql strings.Builder
	var sqlValue strings.Builder
	sqlArgs := make([]any, 0)
	var sqlArgIdx int64
	var sqlArgAppend = func(v any) {
		sqlArgs = append(sqlArgs, v)
		sqlArgIdx++
	}
	// insert into what
	sql.WriteString("insert into ")
	sql.WriteString("\"")
	sql.WriteString(col.TableName())
	sql.WriteString("\"")
	// columns
	sql.WriteString(" ( ")
	kIdx := 0
	for k, v := range col.Mapping() {
		if kIdx != 0 {
			sql.WriteString(",")
			sqlValue.WriteString(",")
		}
		sql.WriteString(k)
		sqlArgAppend(v)
		sqlValue.WriteString("$")
		sqlValue.WriteString(strconv.FormatInt(sqlArgIdx, 10))
		kIdx++
	}
	sql.WriteString(" ) ")
	// where
	sql.WriteString(" values ")
	sql.WriteString(" ( ")
	sql.WriteString(sqlValue.String())
	sql.WriteString(" ) ")
	sql.WriteString(";")
	_, err := db.Exec(context.TODO(), sql.String(), sqlArgs...)
	return err
}
