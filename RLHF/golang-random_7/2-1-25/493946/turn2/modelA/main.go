package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numWorkers = 4
const taskBuffer = 1000

type task struct {
	id   int
	data string
	done chan bool // received by caller when task is done
}

func worker(workerID int, wg *sync.WaitGroup, taskChan chan<- task, quit <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case t, ok := <-taskChan:
			if !ok {
				return // Channel closed, no more tasks
			}
			fmt.Printf("Worker %d: Processing task %d\n", workerID, t.id)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate work
			close(t.done)                                                 // Signal the task is complete
		case <-quit:
			fmt.Printf("Worker %d: Received quit signal, shutting down\n", workerID)
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	quit := make(chan struct{})
	tasks := make(chan task, taskBuffer)
	defer close(tasks) // Close the channel once all tasks are sent
	defer close(quit)  // Close quit channel to shutdown workers

	// Start the worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, tasks, quit)
	}

	// Generate some dummy tasks
	numTasks := 10000
	for i := 0; i < numTasks; i++ {
		done := make(chan bool)
		tasks <- task{i, fmt.Sprintf("Task %d", i), done}

		go func() {
			<-done // Wait for the task to be completed
			fmt.Printf("Task %d completed\n", i)
		}()
	}

	wg.Wait() // Wait for all workers to finish

	fmt.Println("All tasks processed")
}
