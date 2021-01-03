package id63gen

import (
	"testing"
)

func TestNext(t *testing.T) {
	for i := 0; i < 100; i++ {
		println(Next())
	}
}

func BenchmarkNext(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Next()
	}
}
