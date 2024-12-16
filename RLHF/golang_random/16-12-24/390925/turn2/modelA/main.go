package main

import (
	"fmt"
	"sync"
	"time"
)

func squareWorker(numbers <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range numbers {
		results <- n * n
	}
}

func main() {
	start := 1
	end := 100000

	numbers := make(chan int, 100)
	results := make(chan int, 100)

	var wg sync.WaitGroup

	// Starting several workers
	for w := 1; w <= 8; w++ { // e.g. 8 workers
		wg.Add(1)
		go squareWorker(numbers, results, &wg)
	}

	startTime := time.Now()

	// Feeding the input to the workers
	go func() {
		for i := start; i <= end; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	// Gathering results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Output the results
	for result := range results {
		fmt.Println(result)
	}

	elapsed := time.Since(startTime)
	fmt.Printf("Processing took %s\n", elapsed)
}
