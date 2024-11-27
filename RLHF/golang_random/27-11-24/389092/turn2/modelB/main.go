package main

import (
	"fmt"
	"sync"
	"time"
)

const numTasks = 20
const maxWorkers = 5

// workerFunction executes tasks sent through the tasks channel
// It signals that it is done by sending a value on the done channel.
func workerFunction(id int, tasks <-chan int, done chan<- int, wg *sync.WaitGroup) {
	for task := range tasks {
		fmt.Printf("Worker %d started task %d\n", id, task)
		time.Sleep(time.Second * 2) // Simulate work
		fmt.Printf("Worker %d finished task %d\n", id, task)
		done <- id
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	tasks := make(chan int)
	done := make(chan int)

	// Start workers
	numWorkers := maxWorkers
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go workerFunction(i, tasks, done, &wg)
	}

	// Add tasks to the tasks channel
	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}
	close(tasks) // Close the tasks channel to signal no more tasks

	// Collect done signals from workers
	for i := 1; i <= numTasks; i++ {
		<-done
	}

	// Signal to waitgroup that all workers are finished
	close(done)

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All tasks completed.")
}
