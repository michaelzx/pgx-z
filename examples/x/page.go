package x

import (
	"gitee.com/cirsgroup/cirs-one-microlibs/gormlib"
	"gitee.com/cirsgroup/cirs-one-microlibs/pagelib"
	"github.com/michaelzx/pgx-z/pgxz"
	"log"
)

type UserPageRow struct {
	No       string   `db:"no"`
	RealName string   `db:"real_name"`
	WorkNo   string   `db:"work_no"`
	DeptName []string `db:"dept_name"`
}
type UserPageParams struct {
	pgxz.PageParams
	DeptId    int64
	StaffName string
}

func PageByPgxZ() {
	_, err := pgxz.Page[UserPageRow](db,
		"no,real_name,work_no",
		`from "user" where delete_at is null
and (@DeptId=0 or @DeptId = any(dept_ids))
and (@StaffName='' or real_name ilike concat('%',@StaffName,'%'))`,
		"",
		UserPageParams{
			PageParams: pgxz.PageParams{
				PageNum:  1,
				PageSize: 10,
			},
			DeptId: 606633863,
		},
	)
	if err != nil {
		panic(err)
	}
}

const sqlToolsPage = `
select no,real_name,work_no,array(select title from ding_dept where id = any(dept_ids)) as dept_name 
from "user" where delete_at is null
{{if .DeptId}} and @DeptId = any(dept_ids){{end}}
{{if .StaffName}} and real_name ilike concat('%',@StaffName,'%'){{end}}
`

type PageByGormParams struct {
	pagelib.PageParams
	DeptId    int64
	StaffName string
}

func PageByGorm() {
	list := make([]UserPageRow, 0, 0)
	// 不支持limit #,#语法，无法测试~
	_, err := gormlib.Page(gormdb, &list, sqlToolsPage, &PageByGormParams{
		PageParams: pagelib.PageParams{
			PageNum:  1,
			PageSize: 10,
		},
		DeptId: 606633863,
	})
	if err != nil {
		log.Fatalf("%+v", err)
		panic(err)
	}
}
