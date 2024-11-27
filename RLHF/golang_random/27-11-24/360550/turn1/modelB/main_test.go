package main

import (
	"testing"
)

func sum(a, b int) int {
	return a + b
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(100, 200)
	}
}
