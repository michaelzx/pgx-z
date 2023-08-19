package x

import (
	"context"
	"github.com/michaelzx/pgx-z/examples/tmp/col"
	"github.com/michaelzx/pgx-z/examples/tmp/model"
	"github.com/michaelzx/pgx-z/pgxz"
	"time"
)

func UpdateByPgx() {
	db.Exec(context.TODO(), "update ding_dept set real_name=$1 where id=$2",
		"数字化发展部-"+time.Now().Format(time.DateTime),
		835586072)
}
func UpdateByPgxZ() {
	pgxz.Update(db,
		col.DingDept(nil).
			Title("数字化发展部-"+time.Now().Format(time.DateTime)),
		"id=?", 835586072,
	)
}

func UpdateByGorm() {
	gormdb.Model(model.DingDept{}).Where("id=?", 835586072).Updates(map[string]any{
		"title": "数字化发展部-" + time.Now().Format(time.DateTime),
	})
}
