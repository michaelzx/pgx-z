package x

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/michaelzx/pgx-z/examples/internal/col"
	"github.com/michaelzx/pgx-z/examples/internal/model"
	"github.com/michaelzx/pgx-z/pgxz"
)

func GetOneByPgx() {
	rows, _ := db.Query(context.TODO(), "select * from ding_dept where id=$1", 835586072)
	_, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByNameLax[model.DingDept])
	if err != nil {
		panic(err)
	}
}
func GetOneByPgxZ() {
	_, err := pgxz.GetOne[model.DingDept](db, col.DingDept(), "id=?", 835586072)
	if err != nil {
		panic(err)
	}
}

func GetOneByGorm() {
	row := model.DingDept{}
	err := gormdb.Model(model.DingDept{}).Where("id=?", 835586072).First(&row).Error
	if err != nil {
		panic(err)
	}
}
