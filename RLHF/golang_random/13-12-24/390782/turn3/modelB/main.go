package main

import (
	"fmt"
	"sync"
)

func parallelSum(array []int, start, end int, partialSum *int) {
	total := 0
	for i := start; i < end; i++ {
		total += array[i]
	}
	*partialSum += total
}

func main() {
	const numElements = 1000000
	array := make([]int, numElements)
	for i := range array {
		array[i] = i
	}

	var wg sync.WaitGroup
	var totalSum int
	numWorkers := 4
	chunkSize := numElements / numWorkers

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := min((i+1)*chunkSize, numElements)
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			var partialSum int
			parallelSum(array, start, end, &partialSum)
			sync.AddInt64(&totalSum, int64(partialSum))
		}(start, end)
	}

	wg.Wait()
	fmt.Println("Total sum:", totalSum)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
