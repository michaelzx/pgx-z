package x

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/michaelzx/pgx-z/examples/internal/col"
	"github.com/michaelzx/pgx-z/examples/internal/model"
	"github.com/michaelzx/pgx-z/pgxz"
)

func GetAllByPgx() {
	rows, _ := db.Query(context.TODO(), "select * from ding_dept where pid=$1 limit $2", 36091247, 5)
	_, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByNameLax[model.DingDept])
	if err != nil {
		panic(err)
	}
}
func GetAllByPgxZ() {
	rows, err := pgxz.GetAll[model.DingDept](db, col.DingDept(),
		pgxz.Where("pid=?", 36091247),
		pgxz.Limit(5),
	)
	if err != nil {
		panic(err)
	}
	for i, row := range rows {
		fmt.Println(i, row.Id, row.Pid, row.Title)
	}
}

func GetAllByGorm() {
	var rows []model.DingDept
	err := gormdb.Model(model.DingDept{}).Where("pid=?", 36091247).Limit(1).Find(&rows).Error
	if err != nil {
		panic(err)
	}
}
