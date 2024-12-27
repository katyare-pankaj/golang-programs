package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	numGoroutines := 10

	// Add a waitgroup for each goroutine
	wg.Add(numGoroutines)

	// Create goroutines
	for i := 1; i <= numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			doWork(id)
		}(i)
	}

	fmt.Println("Waiting for all goroutines to complete...")
	wg.Wait()
	fmt.Println("All goroutines finished.")
}

func doWork(id int) {
	// Simulate random work duration
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("Goroutine %d completed.\n", id)
}
