package pgxz

import (
	"fmt"
	"github.com/jackc/pgx/v5"
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
func debugPrint(name, sql string, args any) {
	fmt.Printf(">>>>>> pgxz debug %s <<<<<<<<\n", name)
	fmt.Println("[---sql---]")
	fmt.Println(sql)
	fmt.Printf("%T\n", args)
	switch realArgs := args.(type) {
	case []any:
		fmt.Println("[---case([]any)---]")
		for i, arg := range realArgs {
			fmt.Printf("[$%d]=%+v(%T)\n", i+1, arg, arg)
		}
	case pgx.NamedArgs:
		fmt.Println("[---case(pgx.NamedArgs)---]")
		for k, arg := range realArgs {
			fmt.Printf("[@%s]=%+v(%T)\n", k, arg, arg)
		}
	default:
		fmt.Println("[---case(default)---]")
		fmt.Printf("%+v\n", args)
	}
	fmt.Println("**************************")
}
