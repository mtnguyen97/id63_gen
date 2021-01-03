package id63gen

import (
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Next()
	}
}
