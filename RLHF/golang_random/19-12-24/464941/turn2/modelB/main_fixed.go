package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Worker processes tasks from the work channel and adds them to the result channel.
func worker(id int, work <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range work {
		fmt.Printf("Worker %d: processing %d\n", id, num)
		// Simulate work by sleeping for a random duration
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		result <- num * 2
	}
}

func main() {
	numWorkers := 5
	work := make(chan int, 10)   // Buffered channel to limit work in progress
	result := make(chan int, 10) // Buffered channel to hold results
	wg := &sync.WaitGroup{}

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, work, result, wg)
	}

	// Send work to workers
	for i := 1; i <= 20; i++ {
		work <- i
	}

	// Close the work channel to signal workers to stop
	close(work)

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(result) // Close result channel once all workers are done
	}()

	// Collect and print results from result channel
	fmt.Println("Results:")
	for res := range result {
		fmt.Println(res)
	}
}
