package main

import (
	"fmt"
	"sync"
	"time"
)

func processBatch(batch []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, item := range batch {
		fmt.Printf("Processing item %d...\n", item)
		time.Sleep(time.Second) // Simulate some work
	}
}

func main() {
	batches := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9},
	}
	var wg sync.WaitGroup
	wg.Add(len(batches))

	// Start processing each batch concurrently
	for _, batch := range batches {
		go processBatch(batch, &wg)
	}

	// Wait for all batches to be processed before exiting
	wg.Wait()
	fmt.Println("All batches processed.")
}
