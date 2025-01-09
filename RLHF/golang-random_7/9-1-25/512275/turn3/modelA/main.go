package main

import (
	"fmt"
)

func main() {
	const numToPrint = 100

	// Loop to calculate and print the cubes of the first 100 indices
	for i := 0; i < numToPrint; i++ {
		cube := i * i * i
		fmt.Printf("%d^3 = %d\n", i, cube)
	}
}
