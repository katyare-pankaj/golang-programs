package main

import (
	"fmt"
	"sync"
)

// Worker function that processes data and sends it to the output channel
func worker(id int, input <-chan []byte, output chan<- []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range input {
		// Process data in place; for demonstration, we'll just append some text
		processed := append(data, []byte(fmt.Sprintf(" processed by worker %d", id))...)
		output <- processed
	}
}

func main() {
	const (
		numWorkers = 4
		numJobs    = 10
	)

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
			data := []byte(fmt.Sprintf("Some data %d", i))
			input <- data
		}
		close(input) // Close the input channel to signal no more data
	}()

	// Data reader
	for processedData := range output {
		fmt.Println(string(processedData))
	}
}
