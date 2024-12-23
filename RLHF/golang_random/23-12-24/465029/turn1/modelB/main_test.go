package main

import (
	"sync"
	"testing"
)

const (
	numGoroutines = 1000    // Number of goroutines to spawn
	numIterations = 1000000 // Number of iterations for each goroutine
)

func benchmarkIncrement(wg *sync.WaitGroup, counter *int) {
	defer wg.Done()
	for i := 0; i < numIterations; i++ {
		*counter++
	}
}

func BenchmarkIncrementWithWaitGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		counter := 0

		wg.Add(numGoroutines)
		// Start numGoroutines goroutines and pass the WaitGroup pointer
		for j := 0; j < numGoroutines; j++ {
			go benchmarkIncrement(&wg, &counter)
		}
		wg.Wait() // Wait for all goroutines to finish
	}
}
