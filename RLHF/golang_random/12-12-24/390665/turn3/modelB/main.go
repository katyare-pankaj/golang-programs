package main

import (
	"fmt"
	"sync"
)

// CalculateFactorial calculates the factorial of a given number
func calculateFactorial(number int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this Goroutine is done

	if number < 0 {
		result <- 0 // Return 0 for negative numbers
		return
	}

	factorial := 1
	for i := 1; i <= number; i++ {
		factorial *= i
	}
	result <- factorial // Send the result to the result channel
}

func main() {
	var wg sync.WaitGroup
	resultChannel := make(chan int)

	numbers := []int{5, 7, 10} // Numbers for which we want to calculate factorials

	// Start Goroutines to calculate factorials
	for _, number := range numbers {
		wg.Add(1)                                         // Increment the WaitGroup counter
		go calculateFactorial(number, resultChannel, &wg) // Start a Goroutine
	}

	// Wait for all Goroutines to finish
	wg.Wait() // This will block until all Goroutines are done

	// Collect and print results
	close(resultChannel) // Close the channel after all results are sent
	for result := range resultChannel {
		fmt.Printf("Factorial of %d is %d\n", numbers[len(numbers)-1-resultChannel], result)
	}
}
