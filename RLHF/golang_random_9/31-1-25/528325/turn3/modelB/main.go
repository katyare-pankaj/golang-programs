package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// simulateRequest simulates processing a request.
func simulateRequest(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate work by sleeping for a random duration
	workDuration := time.Millisecond * time.Duration(rand.Intn(40)+10)
	time.Sleep(workDuration)

	// Optionally, log the completion of the request
	// fmt.Printf("Request %d processed in %v\n", id, workDuration)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	const (
		minBatchSize = 10
		maxBatchSize = 100
		maxRequests  = 10000
		arrivalRate  = time.Duration(500) * time.Millisecond
	)

	var wg sync.WaitGroup
	totalDuration := time.Duration(0)
	totalBatchCount := 0

	rand.Seed(time.Now().UnixNano())

	requestChan := make(chan int, maxBatchSize)

	go func() {
		for i := 0; i < maxRequests; i++ {
			requestChan <- i
			// Add a random delay between request arrivals
			time.Sleep(time.Duration(rand.Intn(100)) * arrivalRate / 100)
		}
		close(requestChan)
	}()

	for {
		startTime := time.Now()

		// Generate a random batch size between minBatchSize and maxBatchSize
		batchSize := rand.Intn(maxBatchSize-minBatchSize+1) + minBatchSize

		wg.Add(batchSize)
		for i := 0; i < batchSize; i++ {
			// Receive requests from the channel
			reqID, ok := <-requestChan
			if !ok {
				// Channel is closed, no more requests to process
				wg.Wait()
				close(requestChan)
				return
			}

			go simulateRequest(reqID, &wg)
		}

		wg.Wait()
		batchDuration := time.Since(startTime)

		totalDuration += batchDuration
		totalBatchCount++

		// Calculate and print latency and throughput for the batch
		latency := batchDuration / time.Duration(batchSize)
		throughput := float64(batchSize) / batchDuration.Seconds()

		fmt.Printf("Batch %d completed in %v, Latency: %v, Throughput: %.2f requests/sec\n", totalBatchCount, batchDuration, latency, throughput)
	}
}
