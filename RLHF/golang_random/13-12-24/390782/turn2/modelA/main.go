package main

import (
	"fmt"
	"sync"
)

// Function to compute factorial of n
func factorial(n int, wg *sync.WaitGroup, results chan<- int, errors chan<- error) {
	defer wg.Done() // Indicate that this Goroutine is done
	if n < 0 {
		errors <- fmt.Errorf("factorial of negative number %d does not exist", n)
		return
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	results <- result // Send the result back through the channel
}

func main() {
	numbers := []int{5, 4, -1, 7, 6} // List of numbers to compute factorials for
	var wg sync.WaitGroup
	results := make(chan int, len(numbers))  // Results channel
	errors := make(chan error, len(numbers)) // Errors channel

	// Start Goroutines
	for _, n := range numbers {
		wg.Add(1) // Increment the WaitGroup counter
		go factorial(n, &wg, results, errors)
	}

	// Start a Goroutine to wait for all computations to complete
	go func() {
		wg.Wait()      // Wait for all Goroutines to finish
		close(results) // Close the results channel when done
		close(errors)  // Close the errors channel when done
	}()

	// Fetch results and errors
	for i := 0; i < len(numbers); i++ {
		select {
		case res := <-results:
			fmt.Printf("Factorial result: %d\n", res)
		case err := <-errors:
			fmt.Printf("Error: %v\n", err)
		}
	}
}
