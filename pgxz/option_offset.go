package pgxz

import "strconv"

func Limit(v int64) *limitOption {
	return &limitOption{v: v}
}

type limitOption struct {
	v int64
}

func (o *limitOption) ToSql() string {
	return " limit " + strconv.FormatInt(o.v, 10)
}
