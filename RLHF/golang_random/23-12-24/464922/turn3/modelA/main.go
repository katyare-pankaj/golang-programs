package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func task1(wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done
	fmt.Println("Starting task 1...")
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	fmt.Println("Task 1 completed.")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done
	fmt.Println("Starting task 2...")
	time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	fmt.Println("Task 2 completed.")
}

func task3(wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done
	fmt.Println("Starting task 3...")
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Println("Task 3 completed.")
}

func main() {
	var wg sync.WaitGroup // Create a WaitGroup

	// Add three tasks to the WaitGroup
	wg.Add(3)

	// Run task1, task2, and task3 as separate goroutines
	go task1(&wg)
	go task2(&wg)
	go task3(&wg)

	// Wait for all tasks to finish
	fmt.Println("Waiting for tasks to complete...")
	wg.Wait()

	fmt.Println("All tasks completed.")
}
