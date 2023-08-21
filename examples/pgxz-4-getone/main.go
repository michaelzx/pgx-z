package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/michaelzx/pgx-z/examples/internal/col"
	"github.com/michaelzx/pgx-z/examples/internal/model"
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
	pgxz.DEBUG = true
	user, err := pgxz.GetOne[model.Team](db, col.Team(),
		pgxz.Where("no=?", "cjhffdstla5sno0495cg"),
		pgxz.Where("title=?", "4"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	user, err = pgxz.GetOne[model.Team](db, col.Team(),
		pgxz.Where("no=?", "cjhffdstla5sno0495cg"),
		pgxz.Where("title=?", "4"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	domain, err := pgxz.GetOne[model.Domain](db, col.Domain())
	if err != nil {
		panic(err)
	}
	fmt.Println(domain)
}
