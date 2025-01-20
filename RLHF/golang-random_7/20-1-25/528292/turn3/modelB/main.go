package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

const dataSize = 10000000 // 10 million elements

func parallelSort(data []int, numWorkers int) {
	// Create a channel to divide the work into chunks
	chunkChan := make(chan []int)

	// Create a wait group to synchronize the workers
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Launch the workers
	for i := 0; i < numWorkers; i++ {
		go func() {
			for chunk := range chunkChan {
				sort.Ints(chunk)
			}
			wg.Done()
		}()
	}

	// Divide the data into chunks and send them to the workers
	chunkSize := dataSize / numWorkers
	for i := 0; i < dataSize; i += chunkSize {
		end := i + chunkSize
		if end > dataSize {
			end = dataSize
		}
		chunkChan <- data[i:end]
	}

	// Close the channel when all chunks have been sent
	close(chunkChan)

	// Wait for all workers to finish sorting
	wg.Wait()
}

func main() {
	// Generate a large unsorted slice
	data := make([]int, dataSize)
	for i := range data {
		data[i] = rand.Intn(dataSize)
	}

	// Start the timer
	start := time.Now()

	// Parallelize sorting using 4 workers
	parallelSort(data, 4)

	// End the timer and print the execution time
	end := time.Now()
	fmt.Printf("Execution time: %v\n", end.Sub(start))

	// Verify that the data is sorted
	for i := 1; i < dataSize; i++ {
		if data[i] < data[i-1] {
			fmt.Println("Data is not sorted!")
			return