package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Number of goroutines
const numGoroutines = 10000

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(100) * time.Microsecond)
		fmt.Printf("Goroutine %d: %d\n", id, i)
	}
}

func main() {
	var wg sync.WaitGroup

	fmt.Printf("Starting %d goroutines.\n", numGoroutines)
	// Initialize goroutines
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go worker(i, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Printf("All goroutines completed.\n")

	// Output runtime statistics
	runtime.GC()
	fmt.Println("\nRuntime Statistics:")
	runtime.ReadMemStats(&stats)
	fmt.Printf("NumGoroutine: %d\n", stats.NumGoroutine)
	fmt.Printf("NumGoroutine (sched): %d\n", stats.NumGoroutineSched)
}
