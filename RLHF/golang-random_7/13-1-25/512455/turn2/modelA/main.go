package main

import (
	"fmt"
	"sync"
)

// Function that calculates the square of a number and prints it
func calculateAndPrintSquare(number int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this Goroutine is complete
	square := number * number
	fmt.Printf("Square of %d is %d\n", number, square)
}

func main() {
	var wg sync.WaitGroup // Initialize a WaitGroup
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Start a Goroutine for each number in the slice
	wg.Add(len(numbers)) // Indicate the number of Goroutines to wait for
	for _, number := range numbers {
		go calculateAndPrintSquare(number, &wg)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	fmt.Println("All computations have been completed.")
}
