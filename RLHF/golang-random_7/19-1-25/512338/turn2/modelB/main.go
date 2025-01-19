package main

import (
	"fmt"
	"runtime"
	"sync"
)

const (
	numWorkers    = 10
	numIterations = 1000000
)

// WorkerResult stores the result of a worker's task
type WorkerResult struct {
	Index int
	Value int
}

func worker(wg *sync.WaitGroup, input chan int, output chan WorkerResult) {
	defer wg.Done()

	for index := range input {
		// Simulate work
		result := WorkerResult{
			Index: index,
			Value: index * index,
		}

		output <- result
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := &sync.WaitGroup{}
	input := make(chan int, numWorkers)
	output := make(chan WorkerResult, numWorkers)

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(wg, input, output)
	}

	go func() {
		for i := 0; i < numIterations; i++ {
			input <- i
		}
		close(input)
	}()

	results := make([]WorkerResult, numIterations)

	for i := 0; i < numIterations; i++ {
		result := <-output
		results[result.Index] = result
	}

	wg.Wait()

	// Verify results
	for i := 0; i < numIterations; i++ {
		if results[i].Index != i || results[i].Value != i*i {
			fmt.Println("Invalid result:", results[i])
			return
		}
	}

	fmt.Println("All results are correct!")
}
