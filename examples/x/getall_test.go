package x

import "testing"

func BenchmarkGetAll(b *testing.B) {
	benchmarks := []struct {
		name    string
		runFunc func()
	}{
		{"GetAllByPgx", GetAllByPgx},
		{"GetAllByPgxZ", GetAllByPgxZ},
		{"GetAllByGorm", GetAllByGorm},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.runFunc()
			}
		})
	}
}

func TestGetAllByPgxZ(t *testing.T) {
	GetAllByPgxZ()
}
