package main

import (
	"fmt"
)

func main() {
	var slice []int = nil
	var array [3]int

	fmt.Println(fmt.Sprintf("Slice: %v", slice)) // Output: Slice: <nil>
	fmt.Println(fmt.Sprintf("Array: %v", array)) // Output: Array: <nil>

	// Another example with non-nil slice and array
	slice = []int{1, 2, 3}
	array = [3]int{4, 5, 6}

	fmt.Println(fmt.Sprintf("Slice: %v", slice)) // Output: Slice: [1 2 3]
	fmt.Println(fmt.Sprintf("Array: %v", array)) // Output: Array: [4 5 6]
}
