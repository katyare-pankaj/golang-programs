package main

import (
	"fmt"
)

// ClearSensitiveData zeros out all elements in the given slice of bytes to clear sensitive information.
func ClearSensitiveData(sensitiveData []byte) {
	for i := range sensitiveData {
		sensitiveData[i] = 0
	}
}

func main() {
	// Example usage:
	sensitiveData := make([]byte, 1024) // Allocate a large slice for sensitive data

	// Perform some operations on the sensitive data (e.g., encryption, key storage)
	for i := range sensitiveData {
		sensitiveData[i] = 'S' // Simulate sensitive data with the character 'S'
	}

	// Now, we can securely clear the slice
	ClearSensitiveData(sensitiveData)

	// Print the slice to verify that it has been zeroed out
	fmt.Println(sensitiveData) // Output: [0 0 0 ... 0 0]
}
