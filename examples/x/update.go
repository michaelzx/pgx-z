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
		"cjh2vh9u4b8i8e3jnigg")
}
func UpdateByPgxZ() {
	pgxz.Update(db,
		col.Team().
			Title("测试团队-"+time.Now().String()),
		"no=?", "cjh2vh9u4b8i8e3jnigg",
	)
}

func UpdateByGorm() {
	gormdb.Model(model.Team{}).Where("no=?", "cjh2vh9u4b8i8e3jnigg").Updates(map[string]any{
		"title":     "测试团队-" + time.Now().String(),
		"update_at": time.Now(),
	})
}
