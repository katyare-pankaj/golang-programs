package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers    = 5
	taskQueueSize = 10
)

type Task struct {
	id int
}

func worker(id int, tasks <-chan Task, errors chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d started task %d\n", id, task.id)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		if rand.Intn(2) == 0 {
			errors <- fmt.Errorf("worker %d failed task %d", id, task.id)
			continue
		}
		fmt.Printf("Worker %d completed task %d\n", id, task.id)
	}
}

func main() {
	tasks := make(chan Task, taskQueueSize)
	errors := make(chan error, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, errors, &wg)
	}

	go func() {
		for i := 0; i < 15; i++ {
			tasks <- Task{id: i}
		}
		close(tasks)
	}()

	wg.Wait()
	close(errors)

	for err := range errors {
		fmt.Println("Error:", err)
	}

	fmt.Println("All tasks completed.")
}
