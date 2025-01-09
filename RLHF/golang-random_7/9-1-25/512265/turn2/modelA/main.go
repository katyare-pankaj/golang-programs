package main

import (
	"fmt"
	"sync"
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

	var wg sync.WaitGroup
	resultChan := make(chan int, numWorkers)
	chunkCount := (len(data) + chunkSize - 1) / chunkSize // Compute the number of chunks

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, data, chunkSize, chunkCount, resultChan, &wg)
	}

	var sum int
	// Collect results from the channel
	for i := 0; i < chunkCount; i++ {
		sum += <-resultChan
	}

	wg.Wait()
	close(resultChan)

	fmt.Println("Sum:", sum)
}

func worker(id int, data []int, chunkSize int, chunkCount int, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for start := id * chunkSize; start < len(data); start += numWorkers * chunkSize {
		end := min(start+chunkSize, len(data))
		localSum := 0
		for i := start; i < end; i++ {
			localSum += data[i]
		}
		resultChan <- localSum
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
