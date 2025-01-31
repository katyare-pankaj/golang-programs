package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// simulateRequest simulates processing a high-volume request.
func simulateRequest(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate work by sleeping for a random duration between 10 and 50 milliseconds
	workDuration := time.Millisecond * time.Duration(rand.Intn(40)+10)
	time.Sleep(workDuration)

	// Optionally, log the completion of the request
	// fmt.Printf("Request %d processed in %v\n", id, workDuration)
}

func main() {
	// Set GOMAXPROCS to utilize multiple CPUs effectively
	runtime.GOMAXPROCS(runtime.NumCPU())

	const (
		numRequests = 1000 // Total number of requests to process
		batchSize   = 100  // Number of concurrent requests (goroutines)
	)

	var wg sync.WaitGroup
	totalDuration := time.Duration(0)

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numRequests; i += batchSize {
		// Track the start time for this batch
		startTime := time.Now()

		// Read initial GC stats
		var initialMemStats runtime.MemStats
		runtime.ReadMemStats(&initialMemStats)

		// Add the batchSize to the wait group
		wg.Add(batchSize)
		for j := 0; j < batchSize; j++ {
			go simulateRequest(i+j, &wg)
		}

		// Wait for all requests in this batch to finish
		wg.Wait()

		// Calculate and accumulate the duration for this batch
		batchDuration := time.Since(startTime)
		totalDuration += batchDuration

		// Read final GC stats
		var finalMemStats runtime.MemStats
		runtime.ReadMemStats(&finalMemStats)

		// Calculate GC cycles and pause time during this batch
		gcCycles := finalMemStats.NumGC - initialMemStats.NumGC
		gcPauseTime := finalMemStats.PauseTotalNs - initialMemStats.PauseTotalNs

		// Log the batch results
		fmt.Printf("Batch %d completed in %v\n", i/batchSize+1, batchDuration)
		fmt.Printf("GC Cycles: %d, GC Pause Time: %dms\n", gcCycles, gcPauseTime/1e6)
	}

	// Calculate and print overall throughput and latency
	averageDuration := totalDuration / time.Duration(numRequests/batchSize)
	fmt.Printf("\nAverage batch completion time: %v\n", averageDuration)
	fmt.Printf("Overall Throughput: %.2f requests/sec\n", float64(numRequests)/totalDuration.Seconds())
}
