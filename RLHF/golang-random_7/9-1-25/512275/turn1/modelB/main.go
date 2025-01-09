package main

import "fmt"

const (
	arraySize = 10000000
)

func main() {
	data := make([]int, arraySize)
	for i := range data {
		data[i] = i
	}

	// Using pointers to iterate over the array and modify the elements
	for i := range data {
		*(&data[i]) = data[i] * data[i]
	}

	// Print the first 10 elements to verify the modification
	for i := 0; i < 10; i++ {
		fmt.Printf("data[%d] = %d\n", i, data[i])
	}
}
