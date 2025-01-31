package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Worker function simulates a microservices task
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate variable workload by sleeping for a random time
	workDuration := time.Millisecond * time.Duration(rand.Intn(100))
	time.Sleep(workDuration)

	// Log work completion (this log can be removed in production)
	fmt.Printf("Worker %d completed in %v\n", id, workDuration)
}

func main() {
	// Set the number of goroutines as the number of OS threads
	// Normally, you would set this to runtime.NumCPU() for CPU-bound tasks
	runtime.GOMAXPROCS(runtime.NumCPU())

	var (
		numWorkers    = 50
		iterations    = 100
		wg            sync.WaitGroup
		totalDuration time.Duration
	)

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < iterations; i++ {
		// Track start time for latency profiling
		start := time.Now()

		// Get initial GC statistics
		var initialStats runtime.MemStats
		runtime.ReadMemStats(&initialStats)

		// Add workers to the wait group
		wg.Add(numWorkers)
		for j := 0; j < numWorkers; j++ {
			go worker(j, &wg)
		}

		// Wait for all workers to complete
		wg.Wait()

		// Calculate elapsed time for this iteration
		duration := time.Since(start)
		totalDuration += duration

		// Get final GC statistics
		var finalStats runtime.MemStats
		runtime.ReadMemStats(&finalStats)

		// Calculate number of GC cycles and total GC time
		gcCycles := finalStats.NumGC - initialStats.NumGC
		gcTime := finalStats.PauseTotalNs - initialStats.PauseTotalNs

		// Log iteration results
		fmt.Printf("Iteration %d completed in %v\n", i+1, duration)
		fmt.Printf("GC Cycles: %d, GC Time: %dms\n", gcCycles, gcTime/1e6)
	}

	// Calculate average duration
	averageDuration := totalDuration / time.Duration(iterations)

	// Print summary
	fmt.Printf("\nAverage task completion time: %v\n", averageDuration)
	fmt.Printf("Overall Throughput: %.2f tasks/sec\n", float64(numWorkers*iterations)/totalDuration.Seconds())
}
