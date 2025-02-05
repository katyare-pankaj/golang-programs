package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	url         = "http://localhost:8080/benchmark" // Target URL
	numRequests = 100                               // Total number of requests
	concurrency = 10                                // Number of concurrent workers
)

func main() {
	var wg sync.WaitGroup
	responseTimes := make([]time.Duration, numRequests)
	requestsPerWorker := numRequests / concurrency

	// Channel for collecting response times
	results := make(chan time.Duration, numRequests)

	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < requestsPerWorker; j++ {
				start := time.Now()
				if err := sendRequest(); err != nil {
					log.Printf("Request failed: %v", err)
					continue
				}
				duration := time.Since(start)
				results <- duration
			}
		}()
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect all response times
	i := 0
	for duration := range results {
		responseTimes[i] = duration
		i++
	}

	// Calculate and display statistics
	totalDuration := time.Duration(0)
	for _, duration := range responseTimes {
		totalDuration += duration
	}
	averageDuration := totalDuration / time.Duration(len(responseTimes))

	fmt.Printf("Total Requests: %d\n", numRequests)
	fmt.Printf("Concurrency Level: %d\n", concurrency)
	fmt.Printf("Average Response Time: %v\n", averageDuration)
}

func sendRequest() error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}
