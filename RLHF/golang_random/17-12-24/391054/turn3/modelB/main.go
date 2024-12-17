package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	numGoroutines = 16
	numIterations = 100000000
)

func busyWork(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numIterations; i++ {
		_ = i * i * i * i // Simulate CPU-intensive work
	}
}

func main() {
	fmt.Printf("Number of logical CPUs: %d\n", runtime.NumCPU())

	for _, procs := range []int{1, 2, 4, 8, runtime.NumCPU()} {
		runtime.GOMAXPROCS(procs)

		var wg sync.WaitGroup
		startTime := time.Now()

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go busyWork(&wg)
		}

		wg.Wait()
		duration := time.Since(startTime)

		fmt.Printf("GOMAXPROCS: %d, Duration: %s\n", procs, duration)
	}
}
