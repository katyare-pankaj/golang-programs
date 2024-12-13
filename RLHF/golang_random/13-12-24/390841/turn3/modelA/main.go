package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// trackGoroutine runs a simulated task and tracks execution time
func trackGoroutine(wg *sync.WaitGroup, id int) {
	defer wg.Done() // Signal completion when the goroutine finishes

	start := time.Now() // Start time for this goroutine

	// Simulate work (e.g., processing)
	time.Sleep(time.Duration(id%3+1) * 100 * time.Millisecond)

	elapsed := time.Since(start)               // Calculate elapsed time
	activeGoroutines := runtime.NumGoroutine() // Count active goroutines

	fmt.Printf("Goroutine %d finished in %v (Active goroutines: %d)\n", id, elapsed, activeGoroutines)
}

func main() {
	const numGoroutines = 5
	var wg sync.WaitGroup

	// Start multiple goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go trackGoroutine(&wg, i)

		activeGoroutines := runtime.NumGoroutine() // Count active goroutines after launching
		fmt.Printf("Launched goroutine %d (Active goroutines: %d)\n", i, activeGoroutines)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All goroutines have completed execution.")

	// Final count of active goroutines
	finalActiveGoroutines := runtime.NumGoroutine()
	fmt.Printf("Final number of active goroutines: %d\n", finalActiveGoroutines)
}
