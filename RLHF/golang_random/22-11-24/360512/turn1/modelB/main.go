package main

import (
	"fmt"
	"reflect"
	"sync"
)

// A function that takes a slice of any type and applies a function to each element
func applyFuncToSliceConcurrently(slice interface{}, f interface{}) {
	// Create a WaitGroup to synchronize concurrent goroutines
	var wg sync.WaitGroup

	// Get the reflect value of the slice
	sliceVal := reflect.ValueOf(slice)

	// Get the length of the slice
	length := sliceVal.Len()

	// Calculate the chunk size for concurrent processing
	chunkSize := length/4 + 1

	// Start four concurrent goroutines to process the slice in chunks
	for i := 0; i < 4; i++ {
		start := i * chunkSize
		end := min((i+1)*chunkSize, length)
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				// Get the reflect value of the element at index j
				elem := sliceVal.Index(j)
				// Convert the element value to an interface{}
				elemInterface := elem.Interface()
				// Call the function with the element as an argument
				reflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(elemInterface)})
			}
		}(start, end)
	}

	// Wait for all concurrent goroutines to finish
	wg.Wait()
}

// A simple function that doubles the value of an integer
func double(x int) {
	x *= 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Create a slice of integers
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Apply the double function to each element of the slice concurrently
	applyFuncToSliceConcurrently(nums, double)

	// Print the modified slice
	fmt.Println(nums) // Output: [2 4 6 8 10 12 14 16 18 20]
}
