package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	// Number of concurrent API calls
	numWorkers = 10
	// Total number of API calls to make
	numCalls = 1000
)

// Response struct to hold API response details
type Response struct {
	StatusCode int
	Error      error
}

// function to make an API call
func makeAPICall(url string, wg *sync.WaitGroup, ch chan<- Response) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		ch <- Response{Error: err}
		return
	}
	defer resp.Body.Close()
	ch <- Response{StatusCode: resp.StatusCode}
}
func main() {
	startTime := time.Now()
	ch := make(chan Response, numCalls)
	var wg sync.WaitGroup
	// Launch numWorkers goroutines to make API calls
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go makeAPICall("https://api.example.com/data", &wg, ch)
	}
	wg.Wait()
	close(ch)
	var successCount, errorCount int
	var totalTime time.Duration
	for resp := range ch {
		if resp.Error != nil {
			errorCount++
			fmt.Printf("API call failed: %v\n", resp.Error)
		} else {
			successCount++
		}
	}
	totalTime = time.Since(startTime)
	// Calculate throughput
	throughput := float64(successCount) / totalTime.Seconds()
	fmt.Printf("Total API calls made: %d\n", numCalls)
	fmt.Printf("Successful API calls: %d\n", successCount)
	fmt.Printf("Failed API calls: %d\n", errorCount)
	fmt.Printf("Execution time: %s\n", totalTime)
	fmt.Printf("Throughput: %.2f calls/second\n", throughput)
}
