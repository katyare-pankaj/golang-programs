package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate work
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d completed.\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 5

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i+1, &wg)
	}

	fmt.Println("Waiting for workers to complete...")
	wg.Wait()
	fmt.Println("All workers completed.")
}
