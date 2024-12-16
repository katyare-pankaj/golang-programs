package main

import (
	"fmt"
	"sync"
)

func square(num int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()     // Notify the WaitGroup that this goroutine is done
	result := num * num // Calculate the square
	ch <- result        // Send the result to the channel
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	ch := make(chan int, len(numbers)) // Create a buffered channel
	var wg sync.WaitGroup              // Create a WaitGroup

	for _, num := range numbers {
		wg.Add(1)               // Increment the WaitGroup counter
		go square(num, ch, &wg) // Start a goroutine to calculate the square
	}

	// Close the channel in a separate goroutine once all calculations are done
	go func() {
		wg.Wait() // Wait for all goroutines to finish
		close(ch) // Close the channel
	}()

	// Read from the channel until it's closed
	for result := range ch {
		fmt.Println("Square:", result) // Print the result received from the channel
	}
}
