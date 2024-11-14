package main

import (
	"math/rand"
	"sync"
	"time"
)

// Simulate a financial risk assessment function
func simulateRiskAssessment() time.Duration {
	// This function would perform risk assessment logic. For demonstration, we simulate work with random sleep.
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return time.Now().UnixNano()
}

func main() {
	var wg sync.WaitGroup
	numConcurrency := 20 // Number of concurrent requests
	numRequests := 1000  // Total number of requests
	durations := make([]time.Duration, numRequests)

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			start := time.Now().UnixNano()
			_ = simulateRiskAssessment()
			durations[index] = time.Now().UnixNano() - start
		}(i)
	}

	wg.Wait()

	// Capture performance metrics
	totalTime := time.Duration(0)
	for _, duration := range durations {
		totalTime += duration
	}
	avgResponseTime := float64(totalTime) / float64(numRequests)

	println("Load Test Results:")
	println("------------------")
	println("Number of Concurrent Requests:", numConcurrency)
	println("Total Requests:", numRequests)
	println("Average Response Time (ns):", avgResponseTime)
}
