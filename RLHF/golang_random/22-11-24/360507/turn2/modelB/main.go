package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate some work
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println("worker", id, "finished")
}
func main() {
	var wg sync.WaitGroup
	numWorkers := 5
	// Add the number of workers to the wait group
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, &wg)
	}
	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("all workers finished, proceeding to the next stage...")
	// Proceed to the next stage of execution
}
