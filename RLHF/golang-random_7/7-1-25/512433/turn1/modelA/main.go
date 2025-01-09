package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d finished after %d seconds\n", id, time.Duration(id))
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 5

	// Initialize the WaitGroup
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(&wg, i)
	}

	// Wait for all workers to complete
	wg.Wait()

	fmt.Println("All workers have completed.")
}
