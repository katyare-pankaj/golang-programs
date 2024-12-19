package main

import (
	"fmt"
	"sync"
	"time"
)

// task represents a task to be processed
type task struct {
	id       int
	data     string
	priority int // Higher values mean higher priority
}

// worker is a function that processes tasks from a channel
type worker func(chan task)

// runWorkers starts a specified number of workers to process tasks from a channel
func runWorkers(wg *sync.WaitGroup, numWorkers int, tasks chan task, worker worker) {
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(tasks)
		}()
	}
}

// processTask processes a task with a simulated delay
func processTask(tasks chan task) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d with priority %d: %s\n", task.id%10, task.id, task.priority, task.data)
		time.Sleep(100 * time.Millisecond) // Simulated delay
	}
}

func main() {
	// Create a channel for each priority level
	lowPriorityTasks := make(chan task)
	mediumPriorityTasks := make(chan task)
	highPriorityTasks := make(chan task)

	var wg sync.WaitGroup
	numWorkers := 5

	// Start workers for each priority level
	wg.Add(1)
	go runWorkers(&wg, numWorkers, lowPriorityTasks, processTask)
	wg.Add(1)
	go runWorkers(&wg, numWorkers, mediumPriorityTasks, processTask)
	wg.Add(1)
	go runWorkers(&wg, numWorkers, highPriorityTasks, processTask)

	// Generate and send tasks to the appropriate channels based on their priority
	tasks := []task{
		{id: 1, data: "Low priority task 1", priority: 1},
		{id: 2, data: "High priority task 1", priority: 3},
		{id: 3, data: "Medium priority task 1", priority: 2},
		{id: 4, data: "Low priority task 2", priority: 1},
		{id: 5, data: "High priority task 2", priority: 3},
		{id: 6, data: "Medium priority task 2", priority: 2},
		{id: 7, data: "Low priority task 3", priority: 1},
		{id: 8, data: "High priority task 3", priority: 3},
		{id: 9, data: "Medium priority task 3", priority: 2},
	}

	for _, task := range tasks {
		switch task.priority {
		case 1:
			lowPriorityTasks <- task
		case 2:
			mediumPriorityTasks <- task
		case 3:
			highPriorityTasks <- task
		default:
			fmt.Printf("Invalid priority: %d\n", task.priority)
		}
	}

	// Close the task channels to indicate the end of tasks
	close(lowPriorityTasks)
	close(mediumPriorityTasks)
	close(highPriorityTasks)

	wg.Wait()
	fmt.Println("All tasks completed.")
}
