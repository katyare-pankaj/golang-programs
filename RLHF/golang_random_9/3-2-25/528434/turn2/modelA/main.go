package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	numWorkers = 10  // Number of concurrent workers
	numCalls   = 100 // Total number of API calls
)

func makeAPICall(url string, wg *sync.WaitGroup, successCounter *int, mutex *sync.Mutex, ch chan<- time.Duration) {
	defer wg.Done()
	start := time.Now()
	resp, err := http.Get(url)
	elapsed := time.Since(start)

	mutex.Lock()
	defer mutex.Unlock()
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %v, Status Code: %d\n", err, resp.StatusCode)
	} else {
		(*successCounter)++
	}
	ch <- elapsed
}

func main() {
	var wg sync.WaitGroup
	var successCounter int
	mutex := &sync.Mutex{}
	durations := make(chan time.Duration, numCalls)

	apiURL := "https://api.example.com/data" // Replace with the actual API URL
	startTime := time.Now()

	for i := 0; i < numCalls; i++ {
		wg.Add(1)
		go makeAPICall(apiURL, &wg, &successCounter, mutex, durations)
		if i%numWorkers == 0 {
			time.Sleep(50 * time.Millisecond) // Throttle a bit to avoid overwhelming the server
		}
	}

	wg.Wait()
	close(durations)

	totalTime := time.Since(startTime)

	// Calculate average time per call and throughput
	var totalDuration time.Duration
	for d := range durations {
		totalDuration += d
	}
	averageTime := totalDuration / time.Duration(numCalls)
	throughput := float64(successCounter) / totalTime.Seconds()

	fmt.Printf("Total API calls made: %d\n", numCalls)
	fmt.Printf("Successful API calls: %d\n", successCounter)
	fmt.Printf("Execution time: %v\n", totalTime)
	fmt.Printf("Average time per call: %v\n", averageTime)
	fmt.Printf("Throughput: %.2f calls/second\n", throughput)
}
