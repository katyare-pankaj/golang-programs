package main

import (
	"fmt"
)

func main() {
	// Create an array of integers
	var numbers [5]int = {1, 2, 3, 4, 5}

	// Create a slice from the array
	slice := numbers[:]

	// Print the slice's length and capacity
	fmt.Println("Slice Length:", len(slice))
	fmt.Println("Slice Capacity:", cap(slice))

	// Iterate over the slice using a range loop
	for _, num := range slice {
		fmt.Println("Element:", num)
	}

	// Append new elements to the slice
	slice = append(slice, 6, 7, 8)

	// Print the slice again with updated elements
	fmt.Println("Updated Slice Length:", len(slice))
	fmt.Println("Updated Slice Capacity:", cap(slice))

	// Iterate over the updated slice using a range loop
	for _, num := range slice {
		fmt.Println("Updated Element:", num)
	}
}