package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func processTask(task chan int) {
	defer wg.Done()
	for num := range task {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("Processing task %d\n", num)
	}
}

func main() {
	taskChan := make(chan int)
	numTasks := 100

	wg.Add(5) // Assume we have a pool of 5 goroutines
	for i := 0; i < 5; i++ {
		go processTask(taskChan)
	}

	// Sending tasks
	for i := 0; i < numTasks; i++ {
		taskChan <- i
	}
	close(taskChan)

	// Waiting for all goroutines to finish
	wg.Wait()

	fmt.Println("All tasks completed.")
}
