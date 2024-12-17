package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	chunkSize   = 10000 // Define chunk size
	workerCount = 8     // Number of workers
)

// worker function
func worker(id int, jobs <-chan []int, results chan<- int, wg *sync.WaitGroup, errors chan<- error) {
	defer wg.Done()
	for nums := range jobs {
		for _, n := range nums {
			// Simulating processing: squaring the number
			if n < 0 {
				errors <- fmt.Errorf("negative number: %d", n) // Error example
			} else {
				results <- n * n
			}
		}
	}
}

func main() {
	start := 1
	end := 100000
	numbers := make([]int, end-start+1)
	for i := start; i <= end; i++ {
		numbers[i-start] = i
	}

	jobs := make(chan []int, workerCount)
	results := make(chan int, chunkSize)
	errors := make(chan error, workerCount)
	var wg sync.WaitGroup

	// Start worker goroutines
	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg, errors)
	}

	go func() {
		// Handle errors
		for err := range errors {
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	}()

	startTime := time.Now()

	// Chunk the dataset and send to jobs channel
	for i := 0; i < len(numbers); i += chunkSize {
		end := i + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}
		jobs <- numbers[i:end]
	}

	close(jobs) // Close jobs channel after sending all chunks

	// Wait for all workers to finish
	wg.Wait()
	close(results) // Close results channel after all workers are done
	close(errors)  // Close errors channel

	// Collect results
	for res := range results {
		fmt.Println(res) // Output the results (squared numbers)
	}

	elapsed := time.Since(startTime)
	fmt.Printf("Processing took %s\n", elapsed)
}
