package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Time time.Duration
}

func worker(id int, taskCh <-chan Task, wg *sync.WaitGroup) {
	for task := range taskCh {
		fmt.Printf("Worker %d: Starting task %d, processing time: %v\n", id, task.ID, task.Time)
		time.Sleep(task.Time)
		fmt.Printf("Worker %d: Completed task %d\n", id, task.ID)
		wg.Done()
	}
}
func main() {
	const numWorkers = 5
	const numTasks = 20
	taskCh := make(chan Task, numTasks)
	var wg sync.WaitGroup
	// Start the workers
	for i := 0; i < numWorkers; i++ {
		go worker(i+1, taskCh, &wg)
	}
	// Generate tasks with unique IDs and simulated processing times
	for i := 1; i <= numTasks; i++ {
		taskTime := time.Duration(rand.Intn(2000)) * time.Millisecond
		taskCh <- Task{ID: i, Time: taskTime}
	}
	close(taskCh)
	wg.Add(numTasks)
	wg.Wait()
	fmt.Println("All tasks have been completed.")
}
