package pgxz

func GroupBy(sql string) *groupOption {
	return &groupOption{sql: sql}
}

type groupOption struct {
	sql string
}

func (o *groupOption) ToSql() string {
	return " group by" + o.sql
}
