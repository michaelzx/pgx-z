package x

import "testing"

func BenchmarkPage(b *testing.B) {
	benchmarks := []struct {
		name    string
		runFunc func()
	}{
		// {"UpdateByPgx", UpdateByPgx},
		{"PageByPgxZ", PageByPgxZ},
		{"PageByGorm", PageByGorm},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.runFunc()
			}
		})
	}
}
