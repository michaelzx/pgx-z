package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/michaelzx/pgx-z/examples/tmp/col"
	"github.com/michaelzx/pgx-z/examples/tmp/model"
	"github.com/michaelzx/pgx-z/pgxz"
	"os"
)

var db *pgxz.PgDb

func init() {
	connStr := "postgresql://postgres:postgres@127.0.0.1:5432/cirs_gws?"
	pool, err := pgxpool.New(context.TODO(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	db = pgxz.New(pool)
	db.Ping(context.TODO())

}
func main() {
	err := pgxz.Create(db, col.DingDept(&model.DingDept{
		Id:    88888889,
		Pid:   1,
		Title: "测试部门",
	}))
	if err != nil {
		panic(err)
	}
}
