package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	numWorkers = 10                             // Number of concurrent API calls
	numCalls   = 1000                           // Total number of API calls to make
	apiURL     = "https://api.example.com/data" // Replace with your actual API endpoint
)

// Response struct to hold API response details
type Response struct {
	StatusCode int
	Error      error
}

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
		go makeAPICall(apiURL, &wg, ch)
	}

	wg.Wait()
	close(ch)

	var successCount, errorCount int
	var totalTime time.Duration

	// Calculate success rate and execution time
	for resp := range ch {
		if resp.Error != nil {
			errorCount++
		} else {
			successCount++
		}
	}
	totalTime = time.Since(startTime)
	successRate := float64(successCount) / float64(numCalls) * 100

	// Calculate throughput
	throughput := float64(numCalls) / totalTime.Seconds()

	fmt.Printf("Total API calls made: %d\n", numCalls)
	fmt.Printf("Successful API calls: %d (Success Rate: %.2f%%)\n", successCount, successRate)
	fmt.Printf("Failed API calls: %d\n", errorCount)
	fmt.Printf("Execution time: %s\n", totalTime)
	fmt.Printf("Throughput: %.2f calls/second\n", throughput)
}
