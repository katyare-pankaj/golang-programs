package main

import (
	"fmt"
	"testing"
	"time"
)

const (
	numIterations = 1000000
)

func BenchmarkSomeFunction(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Some operation to be benchmarked
	}
}

func MeasureUserExperience(b *testing.B) {
	start := time.Now()
	// User experience metrics collection code
	end := time.Now()
	fmt.Printf("User Experience Duration: %v\n", end.Sub(start))
}

func ExampleUserExperience() {
	b := &testing.B{}
	MeasureUserExperience(b)
	// Output:
	// User Experience Duration: 1.234567s
}
