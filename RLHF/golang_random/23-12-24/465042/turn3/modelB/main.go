package main

import (
	"fmt"
	"sync"
	"time"
)

type task func()

func executeTask(task task, wg *sync.WaitGroup) {
	defer wg.Done()
	task()
}

func main() {
	var wg sync.WaitGroup

	// Define different tasks with their respective operations
	task1 := func() {
		fmt.Println("Starting Task 1...")
		time.Sleep(2 * time.Second)
		fmt.Println("Task 1 completed.")
	}

	task2 := func() {
		fmt.Println("Starting Task 2...")
		time.Sleep(3 * time.Second)
		fmt.Println("Task 2 completed.")
	}

	task3 := func() {
		fmt.Println("Starting Task 3...")
		time.Sleep(1 * time.Second)
		fmt.Println("Task 3 completed.")
	}

	// Add the number of tasks to the WaitGroup
	wg.Add(3)

	// Launch each task in a separate goroutine
	go executeTask(task1, &wg)
	go executeTask(task2, &wg)
	go executeTask(task3, &wg)

	// Wait for all tasks to complete
	wg.Wait()

	fmt.Println("All tasks completed successfully!")
}
