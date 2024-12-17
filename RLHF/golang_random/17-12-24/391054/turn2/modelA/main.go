package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	numGoroutines = 8
	numIterations = 10000000
)

func sumOfSquares(start, end int, wg *sync.WaitGroup, result *int) {
	defer wg.Done()
	sum := 0
	for i := start; i < end; i++ {
		sum += i * i
	}
	*result += sum
}

func main() {
	for _, procs := range []int{1, 2, 4, 8} {
		runtime.GOMAXPROCS(procs) // Set GOMAXPROCS

		var wg sync.WaitGroup
		result := 0
		part := numIterations / numGoroutines

		startTime := time.Now()

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go sumOfSquares(i*part, (i+1)*part, &wg, &result)
		}

		wg.Wait()

		duration := time.Since(startTime)
		fmt.Printf("GOMAXPROCS: %d, Result: %d, Duration: %s\n", procs, result, duration)
	}
}
