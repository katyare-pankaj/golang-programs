package main

import (
	"fmt"
)

// Map applies a function to each element of a slice and returns a new slice of results.
func Map(slice []int, fn func(int) int) []int {
	mapped := make([]int, len(slice))
	for i, v := range slice {
		mapped[i] = fn(v)
	}
	return mapped
}

// Filter returns a new slice containing only the elements that satisfy the predicate function.
func Filter(slice []int, predicate func(int) bool) []int {
	var filtered []int
	for _, v := range slice {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// Reduce applies a function to each element of a slice, accumulating the result.
func Reduce(slice []int, fn func(int, int) int, initial int) int {
	accumulator := initial
	for _, v := range slice {
		accumulator = fn(accumulator, v)
	}
	return accumulator
}

func main() {
	// Sample data: a slice of integers
	numbers := []int{1, 2, 3, 4, 5}

	// Use Map to square each number
	squaredNumbers := Map(numbers, func(x int) int { return x * x })
	fmt.Println("Squared numbers:", squaredNumbers)

	// Use Filter to select even numbers
	evenNumbers := Filter(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Println("Even numbers:", evenNumbers)

	// Use Reduce to sum all numbers
	sumOfNumbers := Reduce(numbers, func(acc, x int) int { return acc + x }, 0)
	fmt.Println("Sum of numbers:", sumOfNumbers)

	// Use Reduce to find the maximum number
	maxNumber := Reduce(numbers, func(acc, x int) int {
		if x > acc {
			return x
		}
		return acc
	}, numbers[0])
	fmt.Println("Max number:", maxNumber)
}
