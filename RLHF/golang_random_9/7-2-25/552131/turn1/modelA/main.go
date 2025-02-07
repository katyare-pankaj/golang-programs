package main

import (
	"fmt"
	"sync"
)

// Map applies a function to each element of a slice and returns a new slice with the results.
func Map[T any, U any](input []T, fn func(T) U) []U {
	output := make([]U, len(input))
	for i, v := range input {
		output[i] = fn(v)
	}
	return output
}

// Filter returns a new slice containing all elements of the original slice that match the predicate.
func Filter[T any](input []T, predicate func(T) bool) []T {
	output := make([]T, 0)
	for _, v := range input {
		if predicate(v) {
			output = append(output, v)
		}
	}
	return output
}

// Reduce reduces a slice into a single value using a reducer function.
func Reduce[T any, U any](input []T, initial U, reducer func(U, T) U) U {
	acc := initial
	for _, v := range input {
		acc = reducer(acc, v)
	}
	return acc
}

// ParallelMap applies a function to each element of a slice and returns a new slice with the results, optimized for large datasets by performing operations in parallel.
func ParallelMap[T any, U any](input []T, fn func(T) U) []U {
	output := make([]U, len(input))
	var wg sync.WaitGroup

	chunkSize := len(input) / 4 // Arbitrary division into four parts – you can adjust this based on performance tests.
	section := func(start, end int) {
		defer wg.Done()
		for i := start; i < end; i++ {
			output[i] = fn(input[i])
		}
	}

	for i := 0; i < 4; i++ { // Create four concurrent sections
		start := i * chunkSize
		end := start + chunkSize
		if i == 3 { // Make sure the last section handles any remainder
			end = len(input)
		}
		wg.Add(1)
		go section(start, end)
	}

	wg.Wait()
	return output
}

func main() {
	// Example usage with numbers.
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Map: Square each number.
	squared := Map(nums, func(n int) int { return n * n })
	fmt.Println("Squared:", squared)

	// Filter: Keep only even numbers.
	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("Evens:", evens)

	// Reduce: Sum all numbers.
	sum := Reduce(nums, 0, func(a, b int) int { return a + b })
	fmt.Println("Sum:", sum)

	// ParallelMap: Square each number (in parallel).
	squaredParallel := ParallelMap(nums, func(n int) int { return n * n })
	fmt.Println("Squared Parallel:", squaredParallel)

	// Even confirms both approach provides similar outputs.
	fmt.Println("Examples equal:", equal(squared, squaredParallel))
}

// Helper function to check equality of two slices of integers.
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
