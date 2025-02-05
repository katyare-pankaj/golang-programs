package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"
)

const (
	url         = "http://localhost:8080/benchmark"
	numRequests = 10000
	concurrency = 100
)

// Result struct to hold response time
type Result struct {
	ResponseTime time.Duration
}

func main() {
	var wg sync.WaitGroup
	results := make(chan Result, numRequests)

	// Start worker goroutines
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go worker(i, &wg, results)
	}

	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Calculate statistics
	totalTime := time.Duration(0)
	var minTime, maxTime time.Duration
	count := 0
	for result := range results {
		totalTime += result.ResponseTime
		if count == 0 || result.ResponseTime < minTime {
			minTime = result.ResponseTime
		}
		if result.ResponseTime > maxTime {
			maxTime = result.ResponseTime
		}
		count++
	}

	averageTime := totalTime / time.Duration(count)
	p90Time := calculateP90(results)

	fmt.Printf("Total Requests: %d\n", count)
	fmt.Printf("Concurrency Level: %d\n", concurrency)
	fmt.Printf("Average Response Time: %v\n", averageTime)
	fmt.Printf("P90 Response Time: %v\n", p90Time)
	fmt.Printf("Min Response Time: %v\n", minTime)
	fmt.Printf("Max Response Time: %v\n", maxTime)
}

func worker(id int, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()
	client := &http.Client{}

	for i := 0; i < numRequests/concurrency; i++ {
		start := time.Now()
		if err := sendRequest(client); err != nil {
			log.Printf("Worker %d: Request failed: %v", id, err)
			continue
		}
		duration := time.Since(start)
		results <- Result{ResponseTime: duration}
	}
}

func sendRequest(client *http.Client) error {
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

// Function to calculate P90
func calculateP90(results <-chan Result) time.Duration {
	var responseTimes []time.Duration

	// Collect all response times in a slice
	for result := range results {
		responseTimes = append(responseTimes, result.ResponseTime)
	}

	// Sort the slice in ascending order
	sort.Slice(responseTimes, func(i, j int) bool {
		return responseTimes[i] < responseTimes[j]
	})

	// Calculate P90 index
	p90Index := int(float64(len(responseTimes)) * 0.9)

	// Return P90 time
	return responseTimes[p90Index]
}
