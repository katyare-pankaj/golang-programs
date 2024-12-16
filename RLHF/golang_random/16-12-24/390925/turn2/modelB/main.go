package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sumArray(array []int) int {
	total := 0
	for _, num := range array {
		total += num
	}
	return total
}

func processChunk(chunk []int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()
	results <- sumArray(chunk)
}

func parallelSum(data []int, numWorkers int) int {
	rand.Seed(time.Now().UnixNano())
	chunkSize := len(data) / numWorkers
	if chunkSize == 0 {
		chunkSize = 1
	}

	results := make(chan<- int)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := min((i+1)*chunkSize, len(data))
		wg.Add(1)
		go processChunk(data[start:end], &wg, results)
	}

	total := 0
	for i := 0; i < numWorkers; i++ {
		total += <-results
	}

	wg.Wait()
	close(results)
	return total
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	data := make([]int, 10000000)
	for i := range data {
		data[i] = rand.Intn(1000)
	}

	startTime := time.Now()
	totalSum := parallelSum(data, 8) // Adjust numWorkers based on CPU cores
	elapsed := time.Since(startTime)

	fmt.Printf("Parallel sum: %d, completed in %s\n", totalSum, elapsed)
}
