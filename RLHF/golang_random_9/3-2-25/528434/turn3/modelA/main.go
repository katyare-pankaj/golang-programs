package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	numWorkers     = 10                             // Number of concurrent API calls
	numCalls       = 100                            // Total number of API calls to make
	apiURL         = "https://api.example.com/data" // Replace with your actual API endpoint
	maxRetries     = 3                              // Maximum number of retries for failed calls
	initialBackoff = 100 * time.Millisecond         // Initial backoff duration
)

type Response struct {
	StatusCode int
	Error      error
}

func makeAPICall(url string, wg *sync.WaitGroup, ch chan<- Response, retries int) {
	defer wg.Done()

	// Exponential backoff logic
	backoff := initialBackoff

	for i := 0; i <= retries; i++ {
		resp, err := http.Get(url)
		if err != nil {
			if i == retries {
				ch <- Response{Error: err}
			} else {
				time.Sleep(backoff)
				backoff *= 2
			}
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			ch <- Response{StatusCode: resp.StatusCode}
			return
		} else if i == retries {
			ch <- Response{StatusCode: resp.StatusCode, Error: fmt.Errorf("HTTP error %d", resp.StatusCode)}
		}

		time.Sleep(backoff)
		backoff *= 2
	}
}

func main() {
	startTime := time.Now()
	ch := make(chan Response, numCalls)
	var wg sync.WaitGroup

	ticker := time.NewTicker(100 * time.Millisecond) // Adjust tick duration as needed for rate limiting
	defer ticker.Stop()

	for i := 0; i < numCalls; i++ {
		<-ticker.C // Rate limit each request
		wg.Add(1)
		go makeAPICall(apiURL, &wg, ch, maxRetries)
	}

	wg.Wait()
	close(ch)

	var successCount, errorCount int

	for resp := range ch {
		if resp.Error != nil {
			errorCount++
			fmt.Printf("API call failed: %v\n", resp.Error)
		} else {
			successCount++
		}
	}

	executionTime := time.Since(startTime)
	successRate := float64(successCount) / float64(numCalls) * 100
	throughput := float64(numCalls) / executionTime.Seconds()

	fmt.Printf("Total API calls made: %d\n", numCalls)
	fmt.Printf("Successful API calls: %d (Success Rate: %.2f%%)\n", successCount, successRate)
	fmt.Printf("Failed API calls: %d\n", errorCount)
	fmt.Printf("Execution time: %s\n", executionTime)
	fmt.Printf("Throughput: %.2f calls/second\n", throughput)
}
