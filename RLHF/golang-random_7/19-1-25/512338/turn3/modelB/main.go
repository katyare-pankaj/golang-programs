package main

import (
	"fmt"
	"sync"
)

const (
	numWorkers = 4
	batchSize  = 1000
	dataSize   = 1000000
)

// OptimizedData struct with padding to prevent false sharing
type OptimizedData struct {
	value int64
	_     [7]int64 // Padding
}

func worker(in <-chan []OptimizedData, out chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	var total int64
	for data := range in {
		for _, d := range data {
			total += d.value
		}
	}
	out <- total
}

func main() {
	// Create input data
	data := make([]OptimizedData, dataSize)
	for i := range data {
		data[i].value = int64(i)
	}

	var wg sync.WaitGroup
	in := make(chan []OptimizedData, numWorkers)
	out := make(chan int64, numWorkers)

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(in, out, &wg)
	}

	// Distribute data to workers in batches
	var currentBatch []OptimizedData
	for i := range data {
		currentBatch = append(currentBatch, data[i])
		if len(currentBatch) == batchSize {
			in <- currentBatch
			currentBatch = nil
		}
	}
	if len(currentBatch) > 0 {
		in <- currentBatch
	}

	close(in)

	// Aggregate results from workers
	var total int64
	for i := 0; i < numWorkers; i++ {
		total += <-out
	}

	wg.Wait()

	// Verify the total
	var expectedTotal int64
	for i := range data {
		expectedTotal += data[i].value
	}
	fmt.Println("Total:", total)
	if total != expectedTotal {
		fmt.Println("Error: Result mismatch")
	} else {
		fmt.Println("Results match")
	}
}
