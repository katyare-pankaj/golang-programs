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
	maxRetries = 3                              // Maximum number of retry attempts for failed calls
	retryDelay = 500 * time.Millisecond         // Delay between retry attempts
)

type Response struct {
	StatusCode int
	Error      error
}

func makeAPICall(url string, wg *sync.WaitGroup, ch chan<- Response) {
	defer wg.Done()
	for retry := 0; retry <= maxRetries; retry++ {
		resp, err := http.Get(url)
		if err != nil {
			ch <- Response{Error: fmt.Errorf("API call failed (attempt %d): %w", retry+1, err)}
			if retry == maxRetries {
				return
			}
			time.Sleep(retryDelay)
			continue
		}
		defer resp.Body.Close()
		ch <- Response{StatusCode: resp.StatusCode}
		return
	}
}
