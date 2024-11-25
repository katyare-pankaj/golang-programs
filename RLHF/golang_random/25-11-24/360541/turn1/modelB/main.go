package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ServiceA generates random numbers and sends them to serviceB for processing
func serviceA(numbers chan<- int, results <-chan int) {
	rand.Seed(time.Now().UnixNano())
	for {
		// Generate a random number and send it to serviceB for processing
		randomNumber := rand.Intn(100)
		numbers <- randomNumber
		fmt.Printf("ServiceA: Sent %d to ServiceB for processing\n", randomNumber)

		// Wait for the result from serviceB
		result := <-results
		fmt.Printf("ServiceA: Received %d from ServiceB\n", result)
	}
}

// ServiceB receives numbers from serviceA, squares them, and sends the results back
func serviceB(numbers <-chan int, results chan<- int) {
	for {
		// Receive a number from serviceA
		number := <-numbers
		fmt.Printf("ServiceB: Received %d from ServiceA\n", number)

		// Square the number and send the result back to serviceA
		result := number * number
		results <- result
		fmt.Printf("ServiceB: Sent %d back to ServiceA\n", result)
	}
}

func main() {
	// Create channels to communicate between services
	numbers := make(chan int)
	results := make(chan int)

	// Start serviceA and serviceB as goroutines
	go serviceA(numbers, results)
	go serviceB(numbers, results)

	// Wait indefinitely to keep the main process running
	select {}
}
