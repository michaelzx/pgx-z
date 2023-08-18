package gen

import "testing"

func TestNew(t *testing.T) {
	New(Config{
		PgDSN:     "postgresql://postgres:postgres@127.0.0.1:5432/cirs_gws?",
		OutputDir: "./tmp",
	})
}
