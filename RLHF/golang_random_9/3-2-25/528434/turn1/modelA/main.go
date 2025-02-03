package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	numCalls       = 100 // Number of API calls
	maxConcurrency = 10  // Maximum number of concurrent API calls
)

func makeAPICall(id int) error {
	// Replace this block with actual API call logic
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non-ok status code: %d", resp.StatusCode)
	}

	// Simulate processing and delay
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("API call %d: Success\n", id)

	return nil
}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup, mu *sync.Mutex, errorChan chan<- error) {
	defer wg.Done()

	for job := range jobs {
		if err := makeAPICall(job); err != nil {
			mu.Lock()
			fmt.Printf("Error in API call %d: %v\n", job, err)
			errorChan <- err
			mu.Unlock()
		}
	}
}

func main() {
	startTime := time.Now()

	var wg sync.WaitGroup
	var mu sync.Mutex

	jobs := make(chan int, numCalls)
	errorChan := make(chan error, numCalls)
	done := make(chan struct{})

	// Fan out workers
	for w := 1; w <= maxConcurrency; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg, &mu, errorChan)
	}

	// Distribute jobs (APIs to call)
	go func() {
		for i := 1; i <= numCalls; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Collect results and errors
	go func() {
		wg.Wait()
		close(errorChan)
		close(done)
	}()

	// Wait for completion and calculate statistics
	errors := 0
	for err := range errorChan {
		if err != nil {
			errors++
		}
	}

	<-done
	elapsed := time.Since(startTime)

	// Metrics
	fmt.Println("Execution Time:", elapsed)
	fmt.Println("Throughput (calls/sec):", float64(numCalls)/elapsed.Seconds())
	fmt.Printf("Total Errors: %d\n", errors)
}
