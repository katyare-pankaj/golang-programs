package main

import (
	"fmt"
	"sync"
)

// Function to calculate the factorial of a number
func factorial(n int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done() // Signal that this Goroutine is done

	// Calculate factorial
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	results <- fact // Send result to the results channel
}

func main() {
	var wg sync.WaitGroup
	results := make(chan int)      // Channel to collect results
	numbers := []int{5, 7, 10, 12} // Numbers to calculate factorials for

	// Start Goroutines to calculate factorials
	for _, number := range numbers {
		wg.Add(1)                          // Increment the WaitGroup counter
		go factorial(number, &wg, results) // Start a Goroutine for each number
	}

	// Start a Goroutine to close the results channel when all calculations are done
	go func() {
		wg.Wait()      // Wait for all Goroutines to finish
		close(results) // Close the results channel
	}()

	// Collect the results
	for result := range results {
		fmt.Println("Factorial:", result)
	}

	fmt.Println("All factorial calculations are done.")
}
