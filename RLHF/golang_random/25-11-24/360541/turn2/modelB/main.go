package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Data string
}

func worker(id int, tasks <-chan Task, results chan<- string) {
	for task := range tasks {
		fmt.Printf("Worker %d started processing task %d: %s\n", id, task.ID, task.Data)
		// Simulate processing time by sleeping for a random duration
		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
		result := fmt.Sprintf("Task %d processed by Worker %d", task.ID, id)
		results <- result
	}
}

func main() {
	// Number of worker goroutines
	numWorkers := runtime.NumCPU()
	tasks := make(chan Task, 100) // Buffered channel to hold tasks
	results := make(chan string)

	wg := &sync.WaitGroup{}
	wg.Add(numWorkers)

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, tasks, results)
		}(i)
	}

	// Generate tasks and send them to the worker channel
	go func() {
		defer close(tasks)
		for i := 1; i <= 20; i++ {
			task := Task{ID: i, Data: fmt.Sprintf("Task data %d", i)}
			tasks <- task
			fmt.Printf("Sent task %d to the queue\n", task.ID)
		}
	}()

	// Receive results from the result channel
	go func() {
		for result := range results {
			fmt.Println(result)
		}
	}()

	wg.Wait()
	close(results)
	fmt.Println("All tasks processed. Exiting...")
}
