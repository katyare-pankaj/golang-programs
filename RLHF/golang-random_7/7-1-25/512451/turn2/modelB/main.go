package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type task struct {
	input     int
	result    int
	completed bool
}

func worker(wg *sync.WaitGroup, tasks chan task) {
	defer wg.Done()
	for task := range tasks {
		// Simulate some work with a random delay
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		task.result = task.input * task.input
		task.completed = true
		fmt.Printf("Task %d completed: %d * %d = %d\n", task.input, task.input, task.input, task.result)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numTasks := 10
	tasks := make(chan task, numTasks)

	// Start workers
	var wg sync.WaitGroup
	numWorkers := 3 // Adjust the number of workers as needed
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(&wg, tasks)
	}

	// Send tasks to the worker pool
	for i := 1; i <= numTasks; i++ {
		tasks <- task{input: i}
	}
	close(tasks)

	// Wait for all tasks to complete
	wg.Wait()

	// Summarize results
	fmt.Println("\nAll tasks completed. Summarizing results:")
	var totalResult int
	for i := 1; i <= numTasks; i++ {
		task := task{input: i}
		tasks <- task
		for !task.completed {
			time.Sleep(10 * time.Millisecond)
		}
		totalResult += task.result
		fmt.Printf("Result for task %d: %d\n", task.input, task.result)
	}
	fmt.Printf("\nTotal sum of results: %d\n", totalResult)
}
