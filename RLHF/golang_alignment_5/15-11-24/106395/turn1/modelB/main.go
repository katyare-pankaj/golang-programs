package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Worker function that processes a chunk of data
func worker(data []int, result chan int) {
	total := 0
	for _, value := range data {
		total += value
	}
	result <- total
}

func main() {
	// Generate a large dataset
	dataSize := 1000000
	dataset := make([]int, dataSize)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < dataSize; i++ {
		dataset[i] = rand.Intn(1000)
	}

	// Split the dataset into smaller chunks for parallel processing
	chunkSize := 100000
	numChunks := dataSize / chunkSize
	chunks := make([][]int, numChunks)
	for i := 0; i < numChunks; i++ {
		start := i * chunkSize
		end := min((i+1)*chunkSize, dataSize)
		chunks[i] = dataset[start:end]
	}

	// Perform parallel processing using worker goroutines
	results := make(chan int, numChunks)
	for _, chunk := range chunks {
		go worker(chunk, results)
	}

	// Aggregate the results from worker goroutines
	totalSum := 0
	for i := 0; i < numChunks; i++ {
		totalSum += <-results
	}

	fmt.Println("Total sum of the dataset:", totalSum)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
