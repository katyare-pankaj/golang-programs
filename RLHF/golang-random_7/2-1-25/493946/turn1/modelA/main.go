package main

import (
	"fmt"
	"sync"
	"time"
)

func processBatch(batch []int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(batch[0]) * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup
	dataset := []int{100, 200, 300, 400, 500}
	batchSize := 2

	// Split the dataset into batches
	for i := 0; i < len(dataset); i += batchSize {
		batch := dataset[i:min(i+batchSize, len(dataset))]
		wg.Add(1)
		go processBatch(batch, &wg)
	}

	// Wait for all batches to be processed
	wg.Wait()

	fmt.Println("All batches processed.")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
