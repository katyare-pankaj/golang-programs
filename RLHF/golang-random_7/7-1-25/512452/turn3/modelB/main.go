package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers = 5   // Number of worker goroutines
	numTasks   = 100 // Number of tasks to process
	taskBuffer = 10  // Buffer size for task channel
)

type Task struct {
	ID    int
	Delay time.Duration
}

func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d: Processing task %d with delay %v\n", id, task.ID, task.Delay)
		time.Sleep(task.Delay)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	tasks := make(chan Task, taskBuffer)
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, tasks, &wg)
	}

	// Generate tasks and send them to the channel
	for i := 1; i <= numTasks; i++ {
		delay := time.Duration(rand.Intn(1000)) * time.Millisecond
		tasks <- Task{ID: i, Delay: delay}
	}

	// Close the task channel to signal workers to stop
	close(tasks)

	// Wait for all workers to complete
	wg.Wait()

	fmt.Println("All tasks completed. Scheduler shutting down.")
}
