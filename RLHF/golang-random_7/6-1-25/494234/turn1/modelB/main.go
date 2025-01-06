// Package main demonstrates slice aliasing and its impact on data integrity.
package main

import (
	"fmt"
)

// SharedData represents a slice of integers shared between services.
type SharedData []int

// ServiceA modifies the shared data.
func ServiceA(data *SharedData) {
	// Modify the first element of the slice
	(*data)[0] = 10
}

// ServiceB reads the shared data.
func ServiceB(data *SharedData) {
	// Read the first element of the slice
	fmt.Println("Service B: Data[0] =", (*data)[0])
}

func main() {
	// Initialize the shared data with a single element
	sharedData := SharedData{0}

	// Pass the address of the shared data to ServiceA
	ServiceA(&sharedData)

	// Pass the address of the shared data to ServiceB
	ServiceB(&sharedData)
}
