package main

import (
	"fmt"
)

func main() {
	// Initialize an array
	numbers := make([]int, 5)
	for i := range numbers {
		numbers[i] = i
	}

	// Create a slice from the array
	sliceA := numbers[:3]

	// Create another slice from the same array, starting from the end of sliceA
	sliceB := numbers[2:5]

	// Modify sliceB
	for i := range sliceB {
		sliceB[i] = sliceB[i] * 2
	}

	// Modifications to sliceB are reflected in the original array and in sliceA
	fmt.Println("Numbers:", numbers)
	fmt.Println("SliceA:", sliceA)
	fmt.Println("SliceB:", sliceB)
}
