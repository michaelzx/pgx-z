package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
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

type UserPageRow struct {
	No       string `db:"no"`
	RealName string `db:"real_name"`
	WorkNo   string `db:"work_no"`
}
type UserPageParams struct {
	pgxz.PageParams
	DeptId    int64
	Domain    string
	StaffName string
}

func main() {
	pgxz.DEBUG = true
	res, err := pgxz.Page[UserPageRow](db,
		"no,real_name,work_no",
		`from "user" where delete_at is null
and (@DeptId=0 or @DeptId = any(dept_ids))
and (@Domain='' or @Domain=any(domain_scope))
and (@StaffName='' or real_name ilike concat('%',@StaffName,''))`,
		"order by work_no desc",
		UserPageParams{
			PageParams: pgxz.PageParams{
				PageNum:  2,
				PageSize: 10,
			},
		},
	)
	if err != nil {
		panic(err)
	}
	for _, row := range res.List {
		fmt.Println(row)
	}

}
