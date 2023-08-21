package pgxz

import "strings"

func Where(sql string, args ...any) *whereOption {
	return &whereOption{sql: sql, args: args}
}

type whereOption struct {
	sql  string
	args []any
}

func (o *whereOption) ToSql() string {
	return o.sql
}

func resolveWheres(wheres ...*whereOption) (string, []any) {
	var sb strings.Builder
	sb.WriteString(" where ")
	args := make([]any, 0)
	for i, where := range wheres {
		if i != 0 {
			sb.WriteString(" and ")
		}
		sb.WriteString("(")
		sb.WriteString(where.sql)
		sb.WriteString(")")
		args = append(args, where.args...)
	}
	return sb.String(), args
}
