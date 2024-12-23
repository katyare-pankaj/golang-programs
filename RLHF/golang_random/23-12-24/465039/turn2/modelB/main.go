package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, numbers <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range numbers {
		fmt.Printf("Worker %d: Processing number %d\n", id, num)
		result := num * num
		time.Sleep(time.Duration(num) * time.Millisecond) // Simulate work
		results <- result
	}
}

func main() {
	numWorkers := 4 // Number of worker goroutines
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numNumbers := len(numbers)

	// Create a WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Create channels for worker communication
	numbersChan := make(chan int, numNumbers)
	resultsChan := make(chan int, numNumbers)

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go worker(i, numbersChan, resultsChan, &wg)
	}

	// Send numbers to workers
	go func() {
		for _, num := range numbers {
			numbersChan <- num
		}
		close(numbersChan)
	}()

	// Collect results from workers
	go func() {
		for i := 0; i < numNumbers; i++ {
			result := <-resultsChan
			fmt.Printf("Result: %d\n", result)
		}
	}()

	// Wait for all workers to complete
	wg.Wait()

	fmt.Println("All workers completed. Program exiting.")
}
