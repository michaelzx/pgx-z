package x

import "testing"

func BenchmarkGetOne(b *testing.B) {
	benchmarks := []struct {
		name    string
		runFunc func()
	}{
		{"GetOneByPgx", GetOneByPgx},
		{"GetOneByPgxZ", GetOneByPgxZ},
		{"GetOneByGorm", GetOneByGorm},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.runFunc()
			}
		})
	}
}
