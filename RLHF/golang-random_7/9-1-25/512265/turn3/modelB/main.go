package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

const numWorkers = 4
const sliceSize = 100000000

func main() {
	data := make([]int, sliceSize)
	for i := 0; i < sliceSize; i++ {
		data[i] = i + 1
	}
	// Use runtime.NumCPU() to get the number of logical CPU cores available
	numWorkers := runtime.NumCPU()
	result := make([]int64, numWorkers)

	// Start worker goroutines
	for w := 0; w < numWorkers; w++ {
		go worker(w, data, &result[w])
	}

	// Wait for all workers to complete
	for _, w := range result {
		fmt.Println(w)
	}
	sum := int64(0)
	for _, r := range result {
		sum += r
	}
	fmt.Println("Final Sum:", sum)
}

func worker(id int, data []int, partialSum *int64) {
	start := id * len(data) / numWorkers
	end := (id + 1) * len(data) / numWorkers
	for _, num := range data[start:end] {
		atomic.AddInt64(partialSum, int64(num))
	}
}
