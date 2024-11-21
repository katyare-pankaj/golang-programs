package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	numTasks := 10
	taskChan := make(chan int, numTasks)
	wg := sync.WaitGroup{}

	// Create worker goroutines
	wg.Add(numTasks)
	for i := 0; i < numTasks; i++ {
		go worker(i, taskChan, &wg)
	}

	// Add tasks to the channel
	for i := 0; i < numTasks; i++ {
		taskChan <- i
	}
	close(taskChan)

	wg.Wait()
	fmt.Println("All tasks completed.")
}

func worker(workerID int, taskChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskChan {
		processTask(workerID, task)
	}
}

func processTask(workerID int, task int) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Worker %d recovered from panic: %v", workerID, r)
		}
	}()

	// Simulate a task that can panic
	time.Sleep(time.Duration(task%3) * time.Second)
	if task%2 == 0 {
		panic("Task failed!")
	}
	fmt.Printf("Worker %d processed task %d\n", workerID, task)
}
