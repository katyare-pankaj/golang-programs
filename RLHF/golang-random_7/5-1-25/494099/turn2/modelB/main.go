package main

import (
	"fmt"
)

func main() {
	// Creating a slice of integers
	// The length of the slice is initially 0, but it can be resized later.
	var slice []int

	// Appending elements to the slice using the append function
	slice = append(slice, 1)
	slice = append(slice, 2, 3, 4, 5)

	// Iterating over the slice using a range loop
	fmt.Println("Elements in the slice:")
	for index, value := range slice {
		fmt.Printf("Index %d: Value %d\n", index, value)
	}

	// Slice literal syntax
	anotherSlice := []string{"apple", "banana", "cherry"}

	// Iterating over the slice literal using a range loop
	fmt.Println("\nElements in anotherSlice:")
	for _, fruit := range anotherSlice {
		fmt.Println(fruit)
	}

	// Slicing a slice
	slicedSlice := slice[1:4] // Gets elements from index 1 to 3 (exclusive)
	fmt.Println("\nElements in slicedSlice:")
	for _, value := range slicedSlice {
		fmt.Println(value)
	}

	// Resizing a slice
	slice = slice[:cap(slice)] // Resizes the slice to its full capacity
	fmt.Println("\nLength of slice after resizing:", len(slice))
}
