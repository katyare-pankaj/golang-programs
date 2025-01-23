package main

import (
	"fmt"
	"sync"
)

const (
	bufSize    = 1024  // Size of the slice buffer
	numWorkers = 4     // Number of worker goroutines
	numJobs    = 10000 // Number of jobs to process
)

// Worker function that processes data and sends it to the output channel.
func worker(id int, input <-chan []byte, output chan<- []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range input {
		// Process data in place
		for i := range data {
			data[i] += 1
		}
		output <- data
	}
}

// Pool for slice allocation
var slicePool = sync.Pool{
	New: func() interface{} {
		return make([]byte, bufSize)
	},
}

func main() {
	input := make(chan []byte)
	output := make(chan []byte)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, input, output, &wg)
	}

	// Start a goroutine to close the output channel once all workers are done
	go func() {
		wg.Wait()
		close(output)
	}()

	// Data generator
	go func() {
		for i := 0; i < numJobs; i++ {
			// Get a slice from the pool
			data := slicePool.Get().([]byte)
			data = data[:0] // Clear the slice
			// Generate data for the slice
			for j := 0; j < bufSize; j++ {
				data = append(data, byte(i%256))
			}
			input <- data
			// Return the slice to the pool
			slicePool.Put(data)
		}
		close(input)
	}()

	// Data reader
	for processedData := range output {
		// Process the processed data
		fmt.Println(len(processedData))
	}
}
