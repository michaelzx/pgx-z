package pgxz

func Select(sql string) *selectOption {
	return &selectOption{sql: sql}
}

type selectOption struct {
	sql  string
	args []any
}

func (o *selectOption) ToSql() string {
	return o.sql
}
