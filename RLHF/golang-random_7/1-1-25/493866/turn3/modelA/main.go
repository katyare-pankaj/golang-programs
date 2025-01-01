package main

import (
	"log"
)

func main() {
	// Test case 1: Logging a nil slice
	var slice1 []int = nil
	log.Printf("Logging nil slice: %v\n", slice1) // Output: Logging nil slice: <nil>

	// Test case 2: Logging a non-nil slice
	slice2 := []int{1, 2, 3}
	log.Printf("Logging non-nil slice: %v\n", slice2) // Output: Logging non-nil slice: [1 2 3]

	// Test case 3: Logging an uninitialized array
	var arr1 [3]int
	log.Printf("Logging uninitialized array: %v\n", arr1) // Output: Logging uninitialized array: [0 0 0]

	// Test case 4: Logging an initialized array
	arr2 := [3]int{4, 5, 6}
	log.Printf("Logging initialized array: %v\n", arr2) // Output: Logging initialized array: [4 5 6]
}
