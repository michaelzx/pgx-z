package x

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/michaelzx/pgx-z/pgxz"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *pgxz.PgDb
var gormdb *gorm.DB

func init() {
	connStr := "postgresql://postgres:postgres@127.0.0.1:5432/cirs_gws?"
	pool, err := pgxpool.New(context.TODO(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	db = pgxz.New(pool)
	db.Ping(context.TODO())

	// https://github.com/go-gorm/postgres
	gormdb, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  connStr,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	gd, err := gormdb.DB()
	if err != nil {
		panic(err)
	}
	gd.Ping()
}
