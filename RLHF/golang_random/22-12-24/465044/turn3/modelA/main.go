package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int, ch chan<- int) {
	defer wg.Done()
	// Simulate varying work load with random sleep time
	time.Sleep(time.Duration(id%5) * time.Second)
	fmt.Printf("Worker %d is done.\n", id)
	ch <- id // Send completion signal to the main goroutine
}

func main() {
	const numWorkers = 10
	var wg sync.WaitGroup
	ch := make(chan<- int, numWorkers) // Buffered channel with buffer size equal to the number of workers

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(&wg, i, ch)
	}

	// Collect completion signals from all workers
	for i := 0; i < numWorkers; i++ {
		signal := <-ch
		fmt.Printf("Worker %d was processed.\n", signal)
	}

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All workers are done.")
}
