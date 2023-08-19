package pgxz

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

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
