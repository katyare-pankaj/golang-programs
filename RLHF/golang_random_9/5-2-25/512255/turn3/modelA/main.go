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
	results := make(chan time.Duration, numRequests)

	// Determine requests per worker
	requestsPerWorker := numRequests / concurrency

	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < requestsPerWorker; j++ {
				startTime := time.Now()
				if err := sendRequest(); err != nil {
					log.Printf("Request failed: %v", err)
					continue
				}
				duration := time.Since(startTime)
				results <- duration
			}
		}(i)
	}

	// Close the results channel once all requests are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect response times
	i := 0
	totalDuration := time.Duration(0)
	for responseTime := range results {
		responseTimes[i] = responseTime
		totalDuration += responseTime
		i++
	}

	averageLatency := totalDuration / time.Duration(len(responseTimes))

	// Output the results
	fmt.Printf("Total Requests: %d\n", numRequests)
	fmt.Printf("Concurrency Level: %d\n", concurrency)
	fmt.Printf("Average Response Time: %v\n", averageLatency)
}

// sendRequest performs an HTTP GET request and returns an error if it fails
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
