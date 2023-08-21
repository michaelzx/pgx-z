package x

import (
	"context"
	"github.com/michaelzx/pgx-z/examples/internal/col"
	"github.com/michaelzx/pgx-z/examples/internal/model"
	"github.com/michaelzx/pgx-z/pgxz"
	"time"
)

func UpdateByPgx() {
	db.Exec(context.TODO(), "update team set title=$1,update_at=$2 where no=$3",
		"测试团队-"+time.Now().String(),
		time.Now(),
		"cjhf63ctla5rq43743o0")
}
func UpdateByPgxZ() {
	pgxz.Update(db,
		col.Team().
			Title("测试团队-"+time.Now().String()),
		pgxz.Where("no=?", "cjhf63ctla5rq43743o0"),
	)
}

func UpdateByGorm() {
	gormdb.Model(model.Team{}).Where("no=?", "cjhf63ctla5rq43743o0").Updates(map[string]any{
		"title":     "测试团队-" + time.Now().String(),
		"update_at": time.Now(),
	})
}
