package gen

import "time"

type table struct {
	NameForDb string    `db:"table_name"`
	NameForGo string    `db:"-"`
	Columns   []column  `db:"-"`
	Imports   []string  `db:"-"`
	NowTime   time.Time `db:"-"`
}
type column struct {
	NameForDb   string  `db:"column_name"`
	NameForGo   string  `db:"-"`
	NameForJson string  `db:"-"`
	TypeForDb   string  `db:"udt_name"`
	TypeForGo   string  `db:"-"`
	IsNullable  bool    `db:"is_nullable"`
	Comment     string  `db:"comment"`
	Default     *string `db:"column_default"`
}
