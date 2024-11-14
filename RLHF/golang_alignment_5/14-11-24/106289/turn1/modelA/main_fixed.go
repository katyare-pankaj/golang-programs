package main

import (
	"math/rand"
	"sync"
	"time"
)

// Simulate a financial risk assessment function
func simulateRiskAssessment() time.Duration {
	// This function would perform risk assessment logic. For demonstration, we simulate work with random sleep.
	sleepDuration := time.Duration(rand.Intn(100)) * time.Millisecond
	time.Sleep(sleepDuration)
	return sleepDuration
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
			start := time.Now()
			_ = simulateRiskAssessment()
			durations[index] = time.Since(start)
		}(i)
	}

	wg.Wait()

	// Capture performance metrics
	totalTime := time.Duration(0)
	for _, duration := range durations {
		totalTime += duration
	}
	avgResponseTime := totalTime / time.Duration(numRequests)

	println("Load Test Results:")
	println("------------------")
	println("Number of Concurrent Requests:", numConcurrency)
	println("Total Requests:", numRequests)
	println("Average Response Time (ns):", avgResponseTime.Nanoseconds())
}
