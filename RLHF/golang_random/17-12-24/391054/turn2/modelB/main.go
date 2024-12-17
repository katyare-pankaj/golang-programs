package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// A simple CPU-bound function that performs busy waiting
func busyWork(id int, duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Goroutine %d starting busy work for %v...\n", id, duration)
	start := time.Now()

	// Busy loop
	for time.Since(start) < duration {
	}

	fmt.Printf("Goroutine %d finished busy work in %v.\n", id, time.Since(start))
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 10
	duration := time.Second * 2 // Each goroutine does 2 seconds of work

	// Set GOMAXPROCS to 1 for sequential execution
	runtime.GOMAXPROCS(1)
	fmt.Println("Running with GOMAXPROCS = 1:")
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go busyWork(i, duration, &wg)
	}
	wg.Wait()
	fmt.Println()

	// Set GOMAXPROCS to runtime.NumCPU() for parallel execution
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Running with GOMAXPROCS =", runtime.GOMAXPROCS(), ":")
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go busyWork(i, duration, &wg)
	}
	wg.Wait()
}
