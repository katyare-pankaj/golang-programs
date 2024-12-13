package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	const numTasks = 10
	var wg sync.WaitGroup

	fmt.Printf("Starting program (Active goroutines: %d)\n", runtime.NumGoroutine())

	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			start := time.Now()

			// Simulate work with random sleep
			duration := time.Duration(rand.Intn(100)) * time.Millisecond
			time.Sleep(duration)

			elapsed := time.Since(start)
			fmt.Printf("Task %d completed in %v (Active goroutines: %d)\n", id, elapsed, runtime.NumGoroutine())
		}(i)

		activeGoroutines := runtime.NumGoroutine()
		fmt.Printf("After launching task %d (Active goroutines: %d)\n", i, activeGoroutines)
	}

	wg.Wait() // Wait for all tasks to finish

	fmt.Printf("All tasks completed (Active goroutines: %d)\n", runtime.NumGoroutine())
}
