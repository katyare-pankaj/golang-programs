package main

import (
	"errors"
	"fmt"
	"sync"
)

// Result represents the result of a computation, potentially with an error.
type Result struct {
	Value interface{}
	Error error
}

// apply applies a function to an argument and returns a Result.
func apply(f func(interface{}) (interface{}, error), arg interface{}) Result {
	value, err := f(arg)
	return Result{Value: value, Error: err}
}

// processData processes data in parallel and returns a slice of Results.
func processData(data []int, f func(int) (interface{}, error)) []Result {
	var wg sync.WaitGroup
	var results []Result

	for _, item := range data {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			result := apply(f, i)
			results = append(results, result)
		}(item)
	}

	wg.Wait()
	return results
}

// double adds one to an integer and returns the result and an error.
func double(x int) (interface{}, error) {
	if x < 0 {
		return nil, errors.New("negative numbers are not allowed")
	}
	return x * 2, nil
}

func main() {
	data := []int{-1, 2, 3, 4, -5}
	results := processData(data, double)

	for _, result := range results {
		if result.Error != nil {
			fmt.Printf("Error: %v\n", result.Error)
		} else {
			fmt.Printf("Value: %v\n", result.Value)
		}
	}
}
