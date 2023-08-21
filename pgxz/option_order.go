package pgxz

func OrderBy(sql string) *orderOption {
	return &orderOption{sql: sql}
}

type orderOption struct {
	sql string
}

func (o *orderOption) ToSql() string {
	return " order by" + o.sql
}
