package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/michaelzx/pgx-z/examples/internal/col"
	"github.com/michaelzx/pgx-z/examples/internal/model"
	"github.com/michaelzx/pgx-z/pgxz"
	"github.com/rs/xid"
	"os"
	"time"
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
	pgxz.DEBUG = true
	err := pgxz.Create(db, col.Team(&model.Team{
		No:       xid.New().String(),
		Title:    "测试部门" + time.Now().String(),
		CreateAt: time.Now(),
		UpdateAt: nil,
		DeleteAt: nil,
	}))
	if err != nil {
		panic(err)
	}
}
