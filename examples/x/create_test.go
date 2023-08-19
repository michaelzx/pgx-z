package x

import "testing"

func BenchmarkCreate(b *testing.B) {
	benchmarks := []struct {
		name    string
		runFunc func()
	}{
		{"CreateByPgx", CreateByPgx},
		{"CreateByPgxZ", CreateByPgxZ},
		{"CreateByGorm", CreateByGorm},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.runFunc()
			}
		})
	}
}
