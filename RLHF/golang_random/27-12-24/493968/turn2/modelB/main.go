package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	ID       int
	Workload int
}

func worker(id int, work chan *Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range work {
		fmt.Printf("Worker %d: Starting task %d with workload %d\n", id, task.ID, task.Workload)
		processTask(task.Workload)
		fmt.Printf("Worker %d: Completed task %d\n", id, task.ID)
	}
}

func processTask(workload int) {
	// Simulate task processing by sleeping for a duration based on the workload.
	// In a real scenario, you would do actual work here.
	duration := time.Duration(workload) * time.Millisecond
	time.Sleep(duration)
}

func main() {
	numWorkers := 4
	tasks := make([]*Task, 0, 20)

	// Generate some tasks with different workloads
	for i := 0; i < 20; i++ {
		workload := rand.Intn(100) + 50
		tasks = append(tasks, &Task{ID: i, Workload: workload})
	}

	work := make(chan *Task)
	var wg sync.WaitGroup

	// Start the worker goroutines
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, work, &wg)
	}

	// Dispatch the tasks to the workers
	for _, task := range tasks {
		work <- task
	}

	close(work)
	wg.Wait()
	fmt.Println("All tasks completed.")
}
