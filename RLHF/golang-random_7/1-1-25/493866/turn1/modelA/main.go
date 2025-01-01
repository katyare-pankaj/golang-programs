package main

import (
	"fmt"
)

func main() {
	var str string
	var strPtr *string = &str
	strPtr = nil

	// Test case 1: Direct string
	fmt.Println(fmt.Sprintf("Direct string: %s", str)) // Output: Direct string:

	// Test case 2: String pointer
	fmt.Println(fmt.Sprintf("String pointer: %s", strPtr)) // Output: String pointer: <nil>

	// Test case 3: Non-nil string pointer
	str = "Hello, world!"
	fmt.Println(fmt.Sprintf("Non-nil string pointer: %s", strPtr)) // Output: Non-nil string pointer:

	// Test case 4: Pointer to non-existent string
	fmt.Println(fmt.Sprintf("Pointer to non-existent string: %s", &str)) // Output: Pointer to non-existent string:
}
