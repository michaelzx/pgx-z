package x

import "testing"

func BenchmarkUpdate(b *testing.B) {
	benchmarks := []struct {
		name    string
		runFunc func()
	}{
		{"UpdateByPgx", UpdateByPgx},
		{"UpdateByPgxZ", UpdateByPgxZ},
		{"UpdateByGorm", UpdateByGorm},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.runFunc()
			}
		})
	}
}
