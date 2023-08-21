package pgxz

import "strconv"

func Offset(v int64) *offsetOption {
	return &offsetOption{v: v}
}

type offsetOption struct {
	v int64
}

func (o *offsetOption) ToSql() string {
	return " offset " + strconv.FormatInt(o.v, 10)
}
