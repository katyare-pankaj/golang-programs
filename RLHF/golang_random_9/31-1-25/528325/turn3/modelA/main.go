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
}

func main() {
	// Set GOMAXPROCS to utilize multiple CPUs effectively
	runtime.GOMAXPROCS(runtime.NumCPU())

	const batchSize = 100 // Number of concurrent requests (goroutines) per batch

	var wg sync.WaitGroup
	var totalRequests int
	var totalDuration time.Duration

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Channel to simulate dynamic incoming requests
	requests := make(chan int, 1000)

	// Simulate generating requests dynamically by sending to the channel
	go func() {
		for i := 0; i < 1234; i++ { // Example: 1234 dynamic requests
			requests <- i
		}
		close(requests)
	}()

	// Process requests in batches
	for batch := 0; true; batch++ {
		batchStartTime := time.Now()

		// Read initial GC stats
		var initialMemStats runtime.MemStats
		runtime.ReadMemStats(&initialMemStats)

		// Collect a batch of requests
		batchRequests := make([]int, 0, batchSize)
		for i := 0; i < batchSize; i++ {
			req, more := <-requests
			if !more {
				break
			}
			batchRequests = append(batchRequests, req)
		}

		// If no more requests, break the loop
		if len(batchRequests) == 0 {
			break
		}

		// Process the batch
		wg.Add(len(batchRequests))
		for _, req := range batchRequests {
			go simulateRequest(req, &wg)
		}

		// Wait for all requests in this batch to finish
		wg.Wait()

		// Calculate and accumulate the duration for this batch
		batchDuration := time.Since(batchStartTime)
		totalDuration += batchDuration
		totalRequests += len(batchRequests)

		// Read final GC stats
		var finalMemStats runtime.MemStats
		runtime.ReadMemStats(&finalMemStats)

		// Calculate GC cycles and pause time during this batch
		gcCycles := finalMemStats.NumGC - initialMemStats.NumGC
		gcPauseTime := finalMemStats.PauseTotalNs - initialMemStats.PauseTotalNs

		// Log the batch results
		fmt.Printf("Batch %d completed in %v with %d requests\n", batch+1, batchDuration, len(batchRequests))
		fmt.Printf("GC Cycles: %d, GC Pause Time: %dms\n", gcCycles, gcPauseTime/1e6)
	}

	// Print overall throughput and latency
	fmt.Printf("\nTotal requests processed: %d\n", totalRequests)
	fmt.Printf("Overall throughput: %.2f requests/sec\n", float64(totalRequests)/totalDuration.Seconds())
}
