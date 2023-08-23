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
	No          string   `db:"no"`
	DingId      string   `db:"ding_id"`
	RealName    string   `db:"real_name"`
	WorkNo      string   `db:"work_no"`
	TeamNo      string   `db:"team_no"`
	TeamTitle   *string  `db:"team_title"`
	Leader      bool     `db:"leader"`
	DomainScope []string `db:"domain_scope"`
	AuthStatus  bool     `db:"auth_status"`
}

type UserPageParams struct {
	pgxz.PageParams
	TeamNo     string
	Domain     string
	StaffName  string
	AuthStatus *bool
}

func main() {
	pgxz.DEBUG = true
	pageParams := UserPageParams{}
	res, err := pgxz.Page[UserPageRow](db,
		`"user".no,ding_id,real_name,work_no,team_no,team.title as team_title,leader,domain_scope,auth_status`,
		`from "user" left join team on team.no="user".team_no
where "user".delete_at is null
and (@TeamNo='' or @TeamNo = team_no)
and (@AuthStatus::boolean is null or @AuthStatus=auth_status)
and (@Domain='' or @Domain=any(domain_scope))
and (@StaffName='' or real_name ilike concat('%',@StaffName,'%'))`,
		"order by leader desc, work_no desc",
		pageParams,
	)
	if err != nil {
		panic(err)
	}
	for _, row := range res.List {
		fmt.Println(row)
	}

}
