package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Initialize variables
	rand.Seed(time.Now().UnixNano())
	numTasks := 100
	maxConcurrentGoroutines := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	metrics := make(map[int]time.Duration)

	// Start tasks
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go func(taskID int) {
			defer wg.Done()
			workDuration := time.Duration(rand.Intn(1000)) * time.Millisecond // Random duration in milliseconds
			time.Sleep(workDuration)
			metrics[taskID] = workDuration
		}(i)
	}

	// Wait for all tasks to complete
	wg.Wait()

	// Analyze metrics
	fmt.Println("Task durations:")
	for taskID, duration := range metrics {
		fmt.Printf("Task %d: %s\n", taskID, duration)
	}

	// Identify imbalances
	maxDuration := max(metrics)
	fmt.Printf("Max task duration: %s\n", maxDuration)

	// Optimize concurrency strategy
	// For simplicity, let's reduce the number of goroutines if the max duration is too high
	if maxDuration > time.Second {
		maxConcurrentGoroutines = maxConcurrentGoroutines / 2
		fmt.Printf("Adjusting max concurrent goroutines to: %d\n", maxConcurrentGoroutines)
	}
}

// Function to find the maximum duration from a map
func max(m map[int]time.Duration) time.Duration {
	var maxDuration time.Duration
	for _, duration := range m {
		if duration > maxDuration {
			maxDuration = duration
		}
	}
	return maxDuration
}
