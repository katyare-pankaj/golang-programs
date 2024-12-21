package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, task chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range task {
		fmt.Printf("Worker %d is processing task %d\n", id, t)
		time.Sleep(time.Second * 1) // Simulate task processing time
	}
}

func main() {
	numWorkers := 5                               // Set the number of workers
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // List of tasks

	var wg sync.WaitGroup
	taskChannel := make(chan int)

	wg.Add(numWorkers)
	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go worker(i, taskChannel, &wg)
	}

	// Push tasks onto the channel
	for _, t := range tasks {
		taskChannel <- t
	}

	close(taskChannel) // Signal no more tasks will be sent

	fmt.Println("Waiting for all workers to finish...")
	wg.Wait()
	fmt.Println("All workers have finished processing tasks.")
}
