package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	// Simulate work
	time.Sleep(time.Duration(1000) * time.Millisecond)
	fmt.Printf("Worker %d completed\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 5

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, &wg)
	}

	fmt.Println("Main goroutine waiting for workers to complete...")
	wg.Wait()
	fmt.Println("All workers completed.")
}
