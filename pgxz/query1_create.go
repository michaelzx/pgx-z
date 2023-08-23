package pgxz

import (
	"context"
	"github.com/jackc/pgx/v5"
	"strconv"
	"strings"
)

func CreateAndReturn[T IModel](db *PgDb, col ICol) (*T, error) {
	sql, sqlArgs := createSqlAndArgs(col)
	sql.WriteString(" RETURNING *;")
	// commit
	if DEBUG {
		debugPrint("create", sql.String(), sqlArgs)
	}
	rows, _ := db.Query(context.TODO(), sql.String(), sqlArgs...)
	return pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByNameLax[T])
}

func Create(db *PgDb, col ICol) error {
	sql, sqlArgs := createSqlAndArgs(col)
	sql.WriteString(";")
	// commit
	if DEBUG {
		debugPrint("create", sql.String(), sqlArgs)
	}
	_, err := db.Exec(context.TODO(), sql.String(), sqlArgs...)
	return err
}

func createSqlAndArgs(col ICol) (*strings.Builder, []any) {
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
	return &sql, sqlArgs
}
