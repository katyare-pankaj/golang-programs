package main

import (
	"fmt"
)

// Add two integers with a private helper function
func Add(a, b int) int {
	return addHelper(a, b)
}

// Private helper function to add two integers
func addHelper(a, b int) int {
	return a + b
}

// API function to get the current year with a private helper function
func GetCurrentYear() int {
	return getCurrentYearHelper()
}

// Private helper function to get the current year
func getCurrentYearHelper() int {
	// Use an actual implementation to get the current year
	return 2023
}

func main() {
	// Call the public functions from the API
	sum := Add(10, 20)
	year := GetCurrentYear()

	fmt.Println("Sum:", sum)   // Output: Sum: 30
	fmt.Println("Year:", year) // Output: Year: 2023

	// Invalid: Cannot access private helper functions directly
	// result := addHelper(100, 200)
	// fmt.Println("Result:", result)
}
