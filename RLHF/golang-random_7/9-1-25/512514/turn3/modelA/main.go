package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a task to be processed.
type Task struct {
	ID   int
	Data string
}

// processTask simulates processing a task.
func processTask(task Task) {
	fmt.Printf("Processing task: ID=%d, Data=%s\n", task.ID, task.Data)
	// Simulate processing time
	time.Sleep(time.Duration(task.ID%100) * time.Millisecond)
}

// producer generates tasks and sends them to the channel.
func producer(taskChannel chan<- Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		task := Task{ID: i, Data: fmt.Sprintf("Data %d", i)}
		taskChannel <- task
		time.Sleep(time.Millisecond * 10)
	}
}

// consumer receives tasks from the channel and processes them.
func consumer(taskChannel <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskChannel {
		processTask(task)
	}
}

func main() {
	const numProducers = 3 // Number of producer goroutines
	const numConsumers = 5 // Number of consumer goroutines
	const bufferSize = 100 // Buffer size for the task channel

	taskChannel := make(chan Task, bufferSize)
	var wg sync.WaitGroup

	// Create producer goroutines
	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go producer(taskChannel, &wg)
	}

	// Create consumer goroutines
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consumer(taskChannel, &wg)
	}

	// Wait for all producers and consumers to finish
	wg.Wait()
	fmt.Println("All tasks processed.")
}
