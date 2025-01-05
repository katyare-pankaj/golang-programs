package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Start 5 worker goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for all worker goroutines to complete
	fmt.Println("Waiting for workers to finish...")
	wg.Wait()

	fmt.Println("All workers have completed.")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started.\n", id)

	// Simulate work with random sleep duration
	sleepDuration := time.Duration(rand.Intn(5)) * time.Second
	time.Sleep(sleepDuration)

	fmt.Printf("Worker %d finished after %s.\n", id, sleepDuration)
}
