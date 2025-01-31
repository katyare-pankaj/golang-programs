package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numGoroutines = 1000    // number of goroutines to run
	numRequests   = 1000000 // total number of requests to handle
	burstSize     = 1000    // size of the burst of requests
)

var wg sync.WaitGroup

func handleRequest(id int, req chan int) {
	for {
		select {
		case r := <-req:
			// Simulate work by doing some arithmetic
			_ = r * r
		default:
			return
		}
	}
}

func main() {
	req := make(chan int, burstSize)

	// Start the goroutines
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go handleRequest(i, req)
	}

	// Generate requests
	start := time.Now()
	for i := 0; i < numRequests; i++ {
		req <- i
	}
	close(req)

	// Wait for all goroutines to finish
	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("Latency: %.2f ms\n", float64(elapsed.Nanoseconds())/float64(numRequests)/1000000)
	fmt.Printf("Throughput: %.2f requests/sec\n", float64(numRequests)/elapsed.Seconds())
}
