package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// trackGoroutine executes a function and tracks its execution time
func trackGoroutine(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	start := time.Now()

	// Simulate work with random sleep
	duration := time.Duration(200+id*50) * time.Millisecond
	time.Sleep(duration)

	elapsed := time.Since(start)
	activeGoroutines := runtime.NumGoroutine()

	fmt.Printf("Goroutine %d finished in %v (Active goroutines: %d)\n", id, elapsed, activeGoroutines)
}

func main() {
	const numGoroutines = 10
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go trackGoroutine(&wg, i)

		activeGoroutines := runtime.NumGoroutine()
		fmt.Printf("Launched goroutine %d (Active goroutines: %d)\n", i, activeGoroutines)
	}

	wg.Wait() // Wait for all goroutines to finish

	fmt.Println("All goroutines have completed execution.")
	rangeActiveGoroutines := runtime.NumGoroutine()
	fmt.Printf("Final number of active goroutines: %d\n", rangeActiveGoroutines)
}
