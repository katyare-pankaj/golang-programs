package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, work <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range work {
		fmt.Printf("Worker %d: processing %d\n", id, num)
		time.Sleep(time.Second)
	}
}

func main() {
	numWorkers := 5
	work := make(chan int, 10) // Buffered channel to limit work in progress
	wg := &sync.WaitGroup{}

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, work, wg)
	}

	// Send work to workers
	for i := 1; i <= 20; i++ {
		work <- i
	}

	// Close the work channel to signal workers to stop
	close(work)

	// Wait for workers to finish
	wg.Wait()
	fmt.Println("All work completed")
}
