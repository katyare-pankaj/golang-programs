package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func processTask(taskId int, task chan int, results chan int) {
	defer wg.Done()
	// Simulate some processing time for each task
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	// Receive task from channel in the given order
	taskIdToProcess := <-task

	// Perform the actual task processing
	result := taskIdToProcess * taskIdToProcess

	// Send the result back using another channel to maintain order
	results <- result
	fmt.Printf("Processed task %d: %d\n", taskId, taskIdToProcess)
}

func main() {
	tasks := []int{1, 2, 3, 4, 5}
	numTasks := len(tasks)

	// Channel to send tasks to worker goroutines in order
	taskCh := make(chan int)

	// Channel to receive results from worker goroutines in order
	resultCh := make(chan int)

	// Start worker goroutines
	wg.Add(numTasks)
	for i := 0; i < numTasks; i++ {
		go processTask(i+1, taskCh, resultCh)
	}

	// Send tasks to the worker goroutines in the correct order
	for _, task := range tasks {
		taskCh <- task
	}

	// Close the task channel to signal the workers that there are no more tasks
	close(taskCh)

	// Receive results from the result channel in the correct order and print them
	for i := 0; i < numTasks; i++ {
		result := <-resultCh
		fmt.Printf("Result for task %d: %d\n", i+1, result)
	}

	wg.Wait()
	close(resultCh)
}
