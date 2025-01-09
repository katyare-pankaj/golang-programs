package main

import (
	"fmt"
)

func main() {
	// Create a large array to hold the squares of the indices
	size := 100000
	squares := make([]int, size)

	// Initialize the array with the squares of the indices
	for i := range squares {
		squares[i] = i * i
	}

	// Print only the first 10 and the last 10 elements
	fmt.Println("First 10 elements:")
	for _, square := range squares[:10] {
		fmt.Printf("%d ", square)
	}
	fmt.Println()

	fmt.Println("Last 10 elements:")
	for _, square := range squares[size-10 : size] {
		fmt.Printf("%d ", square)
	}
	fmt.Println()
}
