package main

import "fmt"

func main() {
	const largeSize = 1_000_000
	const numToPrint = 100

	//Create a slice of integers of length largeSize
	nums := make([]int, largeSize)

	//Initialize the slice with the squares of their indices
	for i := 0; i < largeSize; i++ {
		nums[i] = i * i
	}

	//Print only a subset of the results
	for i := 0; i < numToPrint; i++ {
		fmt.Printf("%d ", nums[i])
	}
}
