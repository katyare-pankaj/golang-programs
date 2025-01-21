package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Attempt to modify the slice during range iteration
	fmt.Println("Original slice:", numbers)

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
		// Appending to a slice could potentially reallocate the underlying array
		// if the slice capacity is exceeded, leading to unexpected behavior.
		if len(numbers) < 10 {
			numbers = append(numbers, i*100) // Change could move to new array
		}
	}

	fmt.Println("Modified slice:", numbers)
	fmt.Println("Note: Modifying a slice during iteration can change the underlying array if appending exceeds capacity.")
}
