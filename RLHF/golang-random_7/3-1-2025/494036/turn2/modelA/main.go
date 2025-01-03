package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a function that simulates a delay
type Task func(wg *sync.WaitGroup)

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task 1 starting...")
	time.Sleep(2 * time.Second) // Simulate delay
	fmt.Println("Task 1 completed.")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task 2 starting...")
	time.Sleep(1 * time.Second) // Simulate delay
	fmt.Println("Task 2 completed.")
}

func task3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task 3 starting...")
	time.Sleep(3 * time.Second) // Simulate delay
	fmt.Println("Task 3 completed.")
}

func main() {
	var wg sync.WaitGroup

	// Add each task to the WaitGroup
	wg.Add(1)
	go task1(&wg)

	wg.Add(1)
	go task2(&wg)

	wg.Add(1)
	go task3(&wg)

	// Wait for all tasks to finish
	wg.Wait()

	fmt.Println("All tasks have completed. Exiting...")
}
