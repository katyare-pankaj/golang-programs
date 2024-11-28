// example.go
package main

import "fmt"

func helloWorld() {
	println("Hello, World!")
}

func add(x, y int) int {
	return x + y
}

var (
	numberOfTries = 10
	flavor        = "chocolate"
)

func processData(data []byte) int {
	byteCount := len(data)
	for i := 0; i < byteCount; i++ {
		data[i] = data[i] * 2
	}
	return byteCount
}

func main() {
	// Call helloWorld function
	helloWorld()

	// Call add function
	x, y := 5, 7
	sum := add(x, y)
	fmt.Printf("Sum of %d and %d is: %d\n", x, y, sum)

	// Use the global variables
	fmt.Printf("Number of tries: %d\n", numberOfTries)
	fmt.Printf("Favorite flavor: %s\n", flavor)

	// Call processData function
	data := []byte{1, 2, 3, 4}
	byteCount := processData(data)
	fmt.Printf("Processed %d bytes. New data: %v\n", byteCount, data)
}
