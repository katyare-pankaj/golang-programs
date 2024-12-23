package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

// Simulate a stateful service with some internal state
type MicroService struct {
	state int
	mu    sync.Mutex
}

func (ms *MicroService) Process(task string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	fmt.Printf("Starting task: %s with state: %d\n", task, ms.state)

	// Simulate some work
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

	// Simulate updating the state
	ms.state++
	fmt.Printf("Task: %s completed with new state: %d\n", task, ms.state)

	return nil
}

func main() {
	// Create a wait group
	var wg sync.WaitGroup

	// Create a micro service instance
	service := &MicroService{state: 0}

	// Add a task to process
	tasks := []string{"task1", "task2", "task3"}
	for _, task := range tasks {
		wg.Add(1)
		go func(task string) {
			defer wg.Done()
			if err := service.Process(task); err != nil {
				fmt.Printf("Error processing task %s: %v\n", task, err)
			}
		}(task)
	}

	// Wait for all tasks to complete
	wg.Wait()

	// Proceed with the next step
	fmt.Println("All tasks are complete. Proceeding with the next step...")
	fmt.Printf("Final state of the service: %d\n", service.state)
}
