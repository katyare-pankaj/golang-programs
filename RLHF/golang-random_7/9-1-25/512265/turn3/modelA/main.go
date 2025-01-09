package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

const (
	numWorkers = 4
	chunkSize  = 1000 // Adjust this size based on your specific use case
)

func main() {
	data := make([]int, 1000000)
	for i := 0; i < len(data); i++ {
		data[i] = i + 1
	}

	var totalSum int64 = 0 // Use int64 to accommodate atomic operations

	// Start worker goroutines
	chunkCount := (len(data) + chunkSize - 1) / chunkSize // Compute the number of chunks

	for i := 0; i < numWorkers; i++ {
		go worker(i, data, chunkSize, chunkCount, &totalSum)
	}

	// Wait for all workers to complete
	for i := 0; i < numWorkers; i++ {
		time.Sleep(time.Millisecond) // Simulate work completion wait, replace with real logic if needed
	}

	fmt.Println("Sum:", int(totalSum))
}

func worker(id int, data []int, chunkSize int, chunkCount int, totalSum *int64) {
	for start := id * chunkSize; start < len(data); start += numWorkers * chunkSize {
		end := min(start+chunkSize, len(data))
		localSum := 0
		for i := start; i < end; i++ {
			localSum += data[i]
		}
		atomic.AddInt64(totalSum, int64(localSum))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
