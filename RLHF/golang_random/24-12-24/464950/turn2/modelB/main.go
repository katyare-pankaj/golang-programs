package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer fmt.Println("Total time:", time.Since(start))

	// Run the benchmark function multiple times to get an average
	const numBenchmarks = 1000000
	for i := 0; i < numBenchmarks; i++ {
		benchmarkDefer()
	}
}

func benchmarkDefer() {
	// Simulate some work being done
	for i := 0; i < 100; i++ {
		_ = i * i
	}

	// Benchmark the overhead of defer
	defer func() {
		// Simulate resource cleanup
		// In reality, this could be anything like closing a file, releasing a lock, or logging.
		_ = time.Sleep(time.Nanosecond * 1)
	}()
}
