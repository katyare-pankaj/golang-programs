package main

import (
	"fmt"
	"unsafe"
)

func cleanUpSlice(s []byte) {
	// Ensure the slice is not nil
	if s == nil {
		return
	}

	// Calculate the length of the underlying array
	sliceLength := len(s)
	cap := cap(s)

	// Zero out the memory of the underlying array
	for i := sliceLength; i < cap; i++ {
		s[i] = 0
	}

	// Set the slice to nil to remove the reference to the array
	s = nil
}

func main() {
	// Create a slice with a capacity larger than its length
	slice := make([]byte, 5, 10)
	copy(slice, "secret")

	// Print the slice
	fmt.Println("Original Slice:", string(slice))

	// Clean up the slice
	cleanUpSlice(slice)

	// Print the slice again to verify that it's empty
	fmt.Println("Cleaned Slice:", string(slice))

	// Access the underlying array to show that it's zeroed out
	array := (*[10]byte)(unsafe.Pointer(&slice[0]))
	fmt.Println("Underlying Array:", string(array[:]))
}
