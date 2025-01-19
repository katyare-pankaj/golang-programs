package main

import (
	"fmt"
	"sync"
)

type DataChunk struct {
	Data []int
}

func processChunk(chunk DataChunk) int {
	sum := 0
	for _, v := range chunk.Data {
		sum += v
	}
	return sum
}

func main() {
	const numChunks = 10
	const chunkSize = 1000000
	const numWorkers = 4

	chunks := make([]DataChunk, numChunks)
	for i := 0; i < numChunks; i++ {
		chunks[i].Data = make([]int, chunkSize)
		for j := 0; j < chunkSize; j++ {
			chunks[i].Data[j] = j
		}
	}

	results := make(chan int, numChunks)
	wg := sync.WaitGroup{}
	semaphore := make(chan struct{}, numWorkers) // Limit number of concurrent workers

	for _, chunk := range chunks {
		wg.Add(1)
		go func(c DataChunk) {
			defer wg.Done()
			semaphore <- struct{}{} // Enter critical section
			result := processChunk(c)
			<-semaphore // Exit critical section
			results <- result
		}(chunk)
	}

	wg.Wait()
	close(results)

	totalSum := 0
	for result := range results {
		totalSum += result
	}

	fmt.Println("Total Sum:", totalSum)
}
