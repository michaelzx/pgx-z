package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/michaelzx/pgx-z/examples/internal/col"
	"github.com/michaelzx/pgx-z/pgxz"
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
	count, err := pgxz.Update(db,
		col.Team().
			Title("update-测试-"+time.Now().Format(time.RFC3339)),
		pgxz.Where("no=?", "cjhf62ctla5rpsqm45og"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("RowsAffected=", count)
}
