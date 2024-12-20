package main

import (
	"errors"
	"fmt"
	"sync"
)

// Processor represents a function that processes data and returns an error.
type Processor func(data int) (int, error)

// processData applies the given processor to each element in the data slice and
// returns a slice of results and a slice of errors.
func processData(data []int, processor Processor) ([]int, []error) {
	var wg sync.WaitGroup
	results := make([]int, len(data))
	errors := make([]error, len(data))

	wg.Add(len(data))
	for i, d := range data {
		go func(i int, d int) {
			defer wg.Done()
			result, err := processor(d)
			results[i] = result
			errors[i] = err
		}(i, d)
	}

	wg.Wait()
	return results, errors
}

// Example processor function that may fail occasionally.
func exampleProcessor(data int) (int, error) {
	if data%2 == 0 {
		return data * data, nil
	}
	return 0, errors.New("failed to process even number")
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	results, errors := processData(data, exampleProcessor)

	// Process the results and errors in a functional way.
	for i, result := range results {
		if errors[i] != nil {
			fmt.Printf("Error processing %d: %v\n", i+1, errors[i])
		} else {
			fmt.Printf("Result %d: %d\n", i+1, result)
		}
	}
}
