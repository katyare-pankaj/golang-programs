package main

import (
	"fmt"
	"sync"
)

func main() {
	// Original slice of numbers
	numbers := []int{1, 2, 3, 4, 5}

	// Create a result slice to store squared numbers
	results := make([]int, len(numbers))

	// Create a channel to send results from goroutines
	resultChannel := make(chan int)

	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup

	// Start goroutines to calculate squares
	for i, num := range numbers {
		wg.Add(1)
		go func(index int, number int) {
			defer wg.Done()
			result := number * number
			resultChannel <- result
			results[index] = result
		}(i, num)
	}

	// Close the channel after all goroutines have finished
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// Collect results from the channel
	for result := range resultChannel {
		fmt.Println(result)
	}
}
