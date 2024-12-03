package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a task to be executed.
type Task func()

// ExecuteConcurrently runs a list of tasks concurrently using Goroutines.
func ExecuteConcurrently(tasks []Task) {
	var wg sync.WaitGroup

	// Helper function to execute a task within a Goroutine.
	runTask := func(task Task) {
		defer wg.Done()
		task()
	}

	// Start Goroutines for each task.
	for _, task := range tasks {
		wg.Add(1)
		go runTask(task)
	}

	// Wait for all Goroutines to finish.
	wg.Wait()
}

func main() {
	tasks := []Task{
		func() {
			fmt.Println("Task 1 is starting.")
			time.Sleep(2 * time.Second)
			fmt.Println("Task 1 is done.")
		},
		func() {
			fmt.Println("Task 2 is starting.")
			time.Sleep(1 * time.Second)
			fmt.Println("Task 2 is done.")
		},
		func() {
			fmt.Println("Task 3 is starting.")
			time.Sleep(3 * time.Second)
			fmt.Println("Task 3 is done.")
		},
	}

	fmt.Println("Executing tasks concurrently...")
	startTime := time.Now()
	ExecuteConcurrently(tasks)
	endTime := time.Now()

	fmt.Println("All tasks completed in:", endTime.Sub(startTime))
}
