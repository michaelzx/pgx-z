package x

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/michaelzx/pgx-z/examples/internal/col"
	"github.com/michaelzx/pgx-z/examples/internal/model"
	"github.com/michaelzx/pgx-z/pgxz"
	"github.com/rs/xid"
	"time"
)

var snowNode *snowflake.Node

func init() {
	var err error
	snowNode, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}
func CreateByPgx() {
	no := xid.New().String()
	_, err := db.Exec(context.TODO(), `insert into team ("no", title)
values ($1,$2);`, no, "CreateByPgx"+time.Now().String())
	if err != nil {
		fmt.Println(no)
		panic(err)
	}
}
func CreateByPgxZ() {
	no := xid.New().String()
	err := pgxz.Create(db, col.Team(&model.Team{
		No:    no,
		Title: "CreateByPgxZ" + time.Now().String(),
	}))
	if err != nil {
		fmt.Println(no)
		panic(err)
	}
}

func CreateByGorm() {
	no := xid.New().String()
	err := gormdb.Create(&model.Team{
		No:    no,
		Title: "CreateByGorm" + time.Now().String(),
	}).Error
	if err != nil {
		fmt.Println(no)
		panic(err)
	}
}
