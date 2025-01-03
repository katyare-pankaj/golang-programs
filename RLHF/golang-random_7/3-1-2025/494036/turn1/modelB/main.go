package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d is starting...\n", id)
	// Simulate work by sleeping for a random duration
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("Task %d is complete!\n", id)
}
func main() {
	var wg sync.WaitGroup
	numTasks := 3
	wg.Add(numTasks)

	// Start the tasks concurrently
	for i := 0; i < numTasks; i++ {
		go task(i+1, &wg)
	}

	// Wait for all tasks to finish
	wg.Wait()
	fmt.Println("All tasks are done.")
}
