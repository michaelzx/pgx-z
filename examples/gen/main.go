package main

import "github.com/michaelzx/pgx-z/gen"

func main() {
	gen.New(gen.Config{
		PgDSN:     "postgresql://postgres:postgres@127.0.0.1:5432/cirs_gws?",
		OutputDir: "./examples/internal",
		PkgName:   "github.com/michaelzx/pgx-z",
	})
}
