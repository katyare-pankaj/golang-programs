package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func processTask(task int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Starting task %d\n", task)
	// Simulate some work
	time.Sleep(2 * time.Second)
	fmt.Printf("Task %d completed\n", task)
}

func monitorProcesses(wg *sync.WaitGroup) {
	for wg.Wait() != nil {
		// Print the number of goroutines that are still running
		fmt.Printf("Number of running goroutines: %d\n", runtime.NumGoroutine())
		time.Sleep(1 * time.Second)
	}
	fmt.Println("All tasks completed")
}

func main() {
	var wg sync.WaitGroup
	numTasks := 5

	wg.Add(numTasks)
	go monitorProcesses(&wg)

	for i := 0; i < numTasks; i++ {
		go processTask(i, &wg)
	}

	wg.Wait()
}
