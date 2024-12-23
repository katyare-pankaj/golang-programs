package main

import (
	"fmt"
	"sync"
	"time"
)

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task 1 is starting.")
	time.Sleep(2 * time.Second)
	fmt.Println("Task 1 is completed.")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task 2 is starting.")
	time.Sleep(1 * time.Second)
	fmt.Println("Task 2 is completed.")
}

func main() {
	var wg sync.WaitGroup

	// Add two to the WaitGroup since we have two tasks
	wg.Add(2)

	// Start both tasks
	go task1(&wg)
	go task2(&wg)

	// Wait for both tasks to finish
	wg.Wait()

	fmt.Println("All tasks are completed. Main function is exiting.")
}
