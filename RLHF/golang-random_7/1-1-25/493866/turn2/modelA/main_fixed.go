package main

import (
	"fmt"
)

func main() {
	// Test case 1: Nil slice
	var slice1 []int = nil
	fmt.Println(fmt.Sprintf("Nil slice: %v", slice1)) // Output: Nil slice: <nil>

	// Test case 2: Non-nil slice
	slice2 := []int{1, 2, 3}
	fmt.Println(fmt.Sprintf("Non-nil slice: %v", slice2)) // Output: Non-nil slice: [1 2 3]

	// Test case 3: Nil array
	var arr1 [3]int
	fmt.Println(fmt.Sprintf("Nil array: %v", arr1)) // Output: Nil array: [0 0 0]

	// Test case 4: Non-nil array
	arr2 := [3]int{4, 5, 6}
	fmt.Println(fmt.Sprintf("Non-nil array: %v", arr2)) // Output: Non-nil array: [4 5 6]
}
