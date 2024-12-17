package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Create a send-only channel for work items
	workCh := make(chan<- string)

	// Create a receive-only channel for results
	resultCh := make(chan string)

	// Worker function that consumes work and produces results
	go func() {
		defer wg.Done()
		for {
			select {
			case task, ok := <-workCh:
				if !ok {
					// Channel is closed, exit worker
					return
				}
				fmt.Println("Working on:", task)
				time.Sleep(time.Second) // Simulate work
				result := fmt.Sprintf("Completed: %s", task)
				resultCh <- result // Send result back
			default:
				time.Sleep(10 * time.Millisecond) // Polling
			}
		}
	}()

	// Consumer function that receives results
	go func() {
		defer wg.Done()
		for result := range resultCh {
			fmt.Println(result)
		}
	}()

	// Send work items
	for i := 0; i < 5; i++ {
		workCh <- fmt.Sprintf("Task %d", i)
	}

	// Close work channel to signal no more work
	close(workCh)

	wg.Wait()
}
