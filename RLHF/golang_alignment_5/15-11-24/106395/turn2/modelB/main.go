package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	dataSize   = 1_000_000 // Large dataset size
	chunkSize  = 10_000    // Size of each data chunk
	numWorkers = 4         // Number of concurrent workers
)

type dataChunk []int

func generateData() dataChunk {
	data := make(dataChunk, dataSize)
	for i := range data {
		data[i] = rand.Intn(1000)
	}
	return data
}

func processChunk(chunk dataChunk, result chan int) {
	var sum int
	for _, value := range chunk {
		sum += value
	}
	result <- sum
}

func main() {
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()

	data := generateData()
	fmt.Println("Data generation completed in", time.Since(startTime))

	numChunks := (dataSize + chunkSize - 1) / chunkSize
	chunks := make([]dataChunk, numChunks)
	for i := range chunks {
		start := i * chunkSize
		end := min((i+1)*chunkSize, dataSize)
		chunks[i] = data[start:end]
	}

	results := make(chan int, numChunks)
	var wg sync.WaitGroup

	for _, chunk := range chunks {
		wg.Add(1)
		go func(chunk dataChunk) {
			defer wg.Done()
			processChunk(chunk, results)
		}(chunk)
	}

	wg.Wait()
	close(results)

	var totalSum int
	for result := range results {
		totalSum += result
	}

	fmt.Println("Data processing completed in", time.Since(startTime))
	fmt.Println("Total sum:", totalSum)

	// Memory cleanup
	data = nil
	for i := range chunks {
		chunks[i] = nil
	}

	runtime.GC()
	fmt.Println("Memory cleaned up and GC completed")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
