package pgxz

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

var DEBUG = false

func wrapErr(err error) error {
	return errors.WithStack(err)
}

type PgDb struct {
	*pgxpool.Pool
}

func New(pool *pgxpool.Pool) *PgDb {
	return &PgDb{
		Pool: pool,
	}
}
func debutPrint(sql string, args []any) {
	fmt.Println("**************************")
	fmt.Println("pgxz debug")
	fmt.Println("**************************")
	fmt.Println(sql)
	for i, arg := range args {
		fmt.Printf("$%d=%+v\n", i+1, arg)
	}
	fmt.Println("**************************")
}
