package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// GoroutineStats holds statistics for a single goroutine
type GoroutineStats struct {
	startTime time.Time
	endTime   time.Time
}

// TrackGoroutineExecution measures the execution time of a goroutine
func TrackGoroutineExecution(w *sync.WaitGroup, f func(), stats *GoroutineStats) {
	defer w.Done()

	stats.startTime = time.Now()
	f()
	stats.endTime = time.Now()
}

// PrintGoroutineStats prints the execution time of the goroutine
func PrintGoroutineStats(stats *GoroutineStats) {
	duration := stats.endTime.Sub(stats.startTime)
	fmt.Printf("Goroutine Execution Time: %s\n", duration)
}

// PrintActiveGoroutines prints the number of active goroutines
func PrintActiveGoroutines() {
	fmt.Printf("Active Goroutines: %d\n", runtime.NumGoroutine())
}

func simulateWork(duration time.Duration) {
	time.Sleep(duration)
}

func main() {
	var wg sync.WaitGroup

	// Start multiple goroutines simulating work
	numGoroutines := 10
	for i := 0; i < numGoroutines; i++ {
		var stats GoroutineStats
		wg.Add(1)
		go TrackGoroutineExecution(&wg, func() {
			simulateWork(time.Duration(rand.Intn(500)) * time.Millisecond)
		}, &stats)
	}

	// Print initial active goroutines
	fmt.Println("Starting...")
	PrintActiveGoroutines()

	// Wait for all goroutines to complete
	wg.Wait()

	// Print final active goroutines
	fmt.Println("\nFinished...")
	PrintActiveGoroutines()

	// Print execution stats for each goroutine
	fmt.Println("\nGoroutine Execution Stats:")
}
