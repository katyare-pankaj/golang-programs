package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"reflect"
)

// SecurelyZeroSliceMemory overwrites the memory allocated for a slice with random bytes.
func SecurelyZeroSliceMemory(slice interface{}) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		log.Fatalf("Invalid argument type. Expecting a slice, got %T", slice)
	}

	// Calculate the number of bytes needed to zero out the slice
	numBytes := s.Len() * s.Elem().Size()

	// Overwrite the slice with random bytes
	_, err := rand.Read(s.Slice(0, numBytes).Bytes())
	if err != nil {
		log.Fatalf("Error reading random bytes: %v", err)
	}
}

func main() {
	// Example usage with sensitive data:
	sensitiveData := []byte("This is sensitive data")

	// Perform some operations on the sensitive data
	fmt.Println(string(sensitiveData))

	// Now, securely clear the sensitive data
	SecurelyZeroSliceMemory(sensitiveData)

	// Verify that the data has been zeroed out
	fmt.Println(string(sensitiveData)) // Output should be a garbled string of random bytes
}
