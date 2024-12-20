package main

import (
	"fmt"
	"sync"
	"time"
)

type TransportationTask struct {
	taskId int
	wg     *sync.WaitGroup
}

func (t *TransportationTask) executeTask() {
	defer t.wg.Done()
	fmt.Printf("Executing task %d...\n", t.taskId)
	// Simulate task execution time
	time.Sleep(time.Duration(t.taskId) * time.Second)
	fmt.Printf("Task %d completed.\n", t.taskId)
}

func main() {
	var wg sync.WaitGroup
	tasks := []TransportationTask{
		{taskId: 1, wg: &wg},
		{taskId: 2, wg: &wg},
		{taskId: 3, wg: &wg},
	}

	wg.Add(len(tasks))

	for _, task := range tasks {
		go task.executeTask()
	}

	fmt.Println("Waiting for all tasks to complete...")
	wg.Wait()
	fmt.Println("All tasks are completed. Starting next phase.")
	// Perform other operations that require all tasks to be complete
}
