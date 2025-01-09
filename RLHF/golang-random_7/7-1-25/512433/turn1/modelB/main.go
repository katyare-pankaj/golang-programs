package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d is starting\n", id)
	// Simulate work being done
	for i := 0; i < 10; i++ {
		fmt.Printf("Worker %d is doing task %d\n", id, i)
	}
	fmt.Printf("Worker %d is done\n", id)
}
func main() {
	var wg sync.WaitGroup
	// Add 3 workers to the waitgroup
	wg.Add(3)
	// Start 3 workers
	for i := 0; i < 3; i++ {
		go worker(i, &wg)
	}
	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers are done")
}
