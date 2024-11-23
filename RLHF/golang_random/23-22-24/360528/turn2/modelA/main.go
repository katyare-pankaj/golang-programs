package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID       int
	Priority int // 1 (high), 2 (medium), 3 (low)
	Function func()
}

func main() {
	// Create channels for each priority level
	highPriority := make(chan Task)
	mediumPriority := make(chan Task)
	lowPriority := make(chan Task)

	// Number of worker goroutines for each priority level
	const highWorkers = 2
	const mediumWorkers = 2
	const lowWorkers = 1

	// Start worker goroutines for each priority
	for i := 0; i < highWorkers; i++ {
		go processTasks(highPriority, "high", &sync.WaitGroup{})
	}

	for i := 0; i < mediumWorkers; i++ {
		go processTasks(mediumPriority, "medium", &sync.WaitGroup{})
	}

	for i := 0; i < lowWorkers; i++ {
		go processTasks(lowPriority, "low", &sync.WaitGroup{})
	}

	// Simulate creating and sending tasks with different priorities
	time.Sleep(1 * time.Second)
	sendTask(highPriority, 1, 1, func() { fmt.Println("High priority task 1") })
	sendTask(mediumPriority, 2, 2, func() { fmt.Println("Medium priority task 2") })
	sendTask(lowPriority, 3, 3, func() { fmt.Println("Low priority task 3") })

	time.Sleep(2 * time.Second)
	sendTask(highPriority, 4, 1, func() { fmt.Println("High priority task 4") })
	sendTask(mediumPriority, 5, 2, func() { fmt.Println("Medium priority task 5") })

	time.Sleep(3 * time.Second)
	sendTask(lowPriority, 6, 3, func() { fmt.Println("Low priority task 6") })

	// Wait for all worker goroutines to finish
	close(highPriority)
	close(mediumPriority)
	close(lowPriority)
}

func sendTask(channel chan Task, id int, priority int, function func()) {
	channel <- Task{ID: id, Priority: priority, Function: function}
}

func processTasks(channel <-chan Task, priorityLevel string, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range channel {
		fmt.Printf("Processing %s priority task %d\n", priorityLevel, task.ID)
		time.Sleep(time.Duration(task.Priority) * time.Second) // Simulate work with time proportional to priority
		task.Function()
	}

	fmt.Println(fmt.Sprintf("%s priority worker finished\n", priorityLevel))
}
