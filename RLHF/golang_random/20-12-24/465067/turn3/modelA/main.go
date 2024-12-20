package main

import (
	"fmt"
	"sync"
)

// Process function that simulates some processing on a number.
func process(num int) int {
	// Simulate work by sleeping for a short period (for demonstration purposes)
	return num * num
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Create a result channel to send results back
	resultChannel := make(chan map[int]int)

	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup

	// Start goroutines to process numbers
	for i, num := range numbers {
		wg.Add(1)
		go func(index int, number int) {
			defer wg.Done()
			result := process(number)
			// Send the result along with its index
			resultChannel <- map[int]int{index: result}
		}(i, num)
	}

	// Close the channel after all goroutines have finished
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// Collect results from the channel
	resultsMap := make(map[int]int)
	for result := range resultChannel {
		for index, value := range result {
			resultsMap[index] = value
		}
	}

	// Print results in order of processing completion
	for index := 0; index < len(numbers); index++ {
		if result, exists := resultsMap[index]; exists {
			fmt.Println(index+1, ":", result)
		} else {
			fmt.Println(index+1, ":", "processing failed")
		}
	}
}
