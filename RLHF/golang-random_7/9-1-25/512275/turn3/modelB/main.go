package main

import "fmt"

func main() {
	const numToPrint = 100
	nums := make([]int, numToPrint)

	// Calculate and print the first numToPrint cube values of indices
	for i := 0; i < numToPrint; i++ {
		nums[i] = i * i * i
		fmt.Printf("%d ", nums[i])
	}
}
