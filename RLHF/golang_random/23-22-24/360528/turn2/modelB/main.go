package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID       int
	Data     string
	Priority int
}

// Define a priority queue using a buffered channel
type PriorityQueue chan Task

func (pq PriorityQueue) enqueue(task Task) {
	pq <- task
}

func (pq PriorityQueue) dequeue() Task {
	return <-pq
}

func main() {
	// Create three priority queues, each with a buffer size of 10
	highPriorityQueue := make(PriorityQueue, 10)
	mediumPriorityQueue := make(PriorityQueue, 10)
	lowPriorityQueue := make(PriorityQueue, 10)

	// Start three worker goroutines for each priority level
	go worker(highPriorityQueue)
	go worker(mediumPriorityQueue)
	go worker(lowPriorityQueue)

	// Generate some sample tasks with different priority levels
	tasks := []Task{
		{ID: 1, Data: "High priority task 1", Priority: 2},
		{ID: 2, Data: "Medium priority task 1", Priority: 1},
		{ID: 3, Data: "Low priority task 1", Priority: 0},
		{ID: 4, Data: "High priority task 2", Priority: 2},
		{ID: 5, Data: "Medium priority task 2", Priority: 1},
		{ID: 6, Data: "Low priority task 2", Priority: 0},
	}

	// Distribute tasks to the appropriate priority queues
	for _, task := range tasks {
		switch task.Priority {
		case 2:
			highPriorityQueue.enqueue(task)
		case 1:
			mediumPriorityQueue.enqueue(task)
		case 0:
			lowPriorityQueue.enqueue(task)
		}
	}

	// Wait for all tasks to be completed
	time.Sleep(5 * time.Second)
}

func worker(queue PriorityQueue) {
	for {
		task := queue.dequeue()
		fmt.Printf("Processing Task %d: %s (Priority: %d)\n", task.ID, task.Data, task.Priority)
		time.Sleep(1 * time.Second) // Simulate task processing time
	}
}
