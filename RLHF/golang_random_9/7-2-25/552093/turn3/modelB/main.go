package main

import (
	"fmt"
	"sync"
	"time"
)

func task1(wg *sync.WaitGroup) {
	fmt.Println("Task 1 started...")
	time.Sleep(3 * time.Second) // Simulate longer task
	fmt.Println("Task 1 completed.")
	wg.Done()
}

func task2(wg *sync.WaitGroup) {
	fmt.Println("Task 2 started...")
	time.Sleep(2 * time.Second) // Simulate shorter task
	fmt.Println("Task 2 completed.")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go task1(&wg)
	go task2(&wg)

	wg.Wait() // Main goroutine will wait here until both tasks finish
	fmt.Println("All tasks completed.")
}
