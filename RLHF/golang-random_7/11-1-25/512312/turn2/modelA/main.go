package main

import (
	"fmt"
)

func main() {
	// Original slice of integers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// New slice to store even numbers
	evenNumbers := []int{}

	// Use a range loop to iterate through the original slice
	for _, num := range numbers {
		// Check if the number is even
		if num%2 == 0 {
			// Append the even number to the new slice
			evenNumbers = append(evenNumbers, num)
		}
	}

	// Print the resulting slice of even numbers
	fmt.Println("Even numbers:", evenNumbers)
}
