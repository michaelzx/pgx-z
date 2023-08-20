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
	_ = pgxz.Update(db,
		col.User(nil).
			RealName("zxp"+time.Now().Format(time.DateTime)),
		"no=?", "cjdgoostla5k7f1kjc8g",
	)
}
