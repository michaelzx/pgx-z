package pgxz

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"reflect"
)

type PageResult[T any] struct {
	Pagination
	List []T
}

type UserPageParams struct {
	DeptId int64
}

func structToNamedArgs(obj any) pgx.NamedArgs {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(pgx.NamedArgs)
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
func (p *PageResult[T]) doQuery(db *PgDb, selectSql, filterSql, orderSql string, pageParams IPageParams) error {
	var err error
	countSql := "select count(*) " + filterSql
	namedArgs := structToNamedArgs(pageParams)
	if DEBUG {
		debugPrint("page.count", countSql, namedArgs)
	}
	var total int64
	err = db.QueryRow(context.TODO(), countSql, namedArgs).Scan(&total)
	if err != nil {
		return wrapErr(err)
	}
	// 计算分页参数
	p.Pagination.Compute(total)
	// 获取分页数据
	pageSql := fmt.Sprintf(`select %s %s %s offset %d limit %d`,
		selectSql, filterSql, orderSql,
		p.GetSkipRows(), p.PageSize,
	)
	if DEBUG {
		debugPrint("page.list", pageSql, namedArgs)
	}
	rows, _ := db.Query(context.TODO(), pageSql, namedArgs)
	list, err := pgx.CollectRows[T](rows, pgx.RowToStructByNameLax[T])
	if err != nil {
		return err
	}
	p.List = list
	return nil
}
