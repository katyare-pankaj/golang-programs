package main

import (
	"fmt"
	"sync"
	"time"
)

func processBatch(batch []int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	for _, item := range batch {
		fmt.Printf("Processing item %d in batch...\n", item)
		time.Sleep(time.Second * 2) // Simulate some work
	}
	results <- "Batch completed"
}

func main() {
	batches := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8},
		{9, 10, 11, 12, 13, 14},
	}

	var wg sync.WaitGroup
	results := make(chan string, len(batches))

	wg.Add(len(batches))

	// Start processing each batch concurrently
	for _, batch := range batches {
		go processBatch(batch, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Wait for all batches to be processed, and print results
	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("All batches processed.")
}
