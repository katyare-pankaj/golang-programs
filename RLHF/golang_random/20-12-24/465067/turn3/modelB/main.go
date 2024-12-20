package main

import (
	"fmt"
	"sync"
)

func processNumber(num int, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done()
	// Simulate processing time
	result := num * num
	resultChan <- result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	resultChan := make(chan int)

	// Start goroutines for processing
	for _, num := range numbers {
		wg.Add(1)
		go processNumber(num, &wg, resultChan)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Receive and print results in order
	for result := range resultChan {
		fmt.Println(result)
	}
}
