package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task represents a unit of work to be processed concurrently.
type Task struct {
	ID          int
	Data        string
	ProcessTime time.Duration
}

// Worker represents a goroutine that processes tasks from a channel.
type Worker struct {
	ID          int
	TaskChannel <-chan Task
	wg          *sync.WaitGroup
}

// NewWorker creates a new Worker goroutine.
func NewWorker(id int, taskChannel <-chan Task, wg *sync.WaitGroup) *Worker {
	worker := &Worker{
		ID:          id,
		TaskChannel: taskChannel,
		wg:          wg,
	}
	wg.Add(1)
	go worker.run()
	return worker
}

// run starts the Worker goroutine, which processes tasks from the channel until it is closed.
func (w *Worker) run() {
	defer w.wg.Done()
	for task := range w.TaskChannel {
		fmt.Printf("Worker %d started processing task %d\n", w.ID, task.ID)
		time.Sleep(task.ProcessTime)
		fmt.Printf("Worker %d finished processing task %d\n", w.ID, task.ID)
	}
}

// manageTasks concurrently processes tasks from multiple data sources using goroutines and channels.
func manageTasks(numWorkers int, taskChannels ...<-chan Task) {
	var wg sync.WaitGroup

	// Create workers
	for workerID := 1; workerID <= numWorkers; workerID++ {
		NewWorker(workerID, mergeChannels(taskChannels...), &wg)
	}

	// Wait for all workers to finish
	wg.Wait()
}

// mergeChannels merges multiple channels into a single channel that delivers tasks from all sources.
func mergeChannels(channels ...<-chan Task) <-chan Task {
	mergedChannel := make(chan Task)
	var wg sync.WaitGroup

	// Start a goroutine for each input channel to forward tasks to the merged channel
	for _, channel := range channels {
		wg.Add(1)
		go func(c <-chan Task) {
			defer wg.Done()
			for task := range c {
				mergedChannel <- task
			}
		}(channel)
	}

	// Go routine to close the merged channel when all input channels are closed
	go func() {
		wg.Wait()
		close(mergedChannel)
	}()

	return mergedChannel
}

func main() {
	// Simulate multiple data sources providing tasks
	dataSource1 := make(chan Task, 10)
	dataSource2 := make(chan Task, 10)
	dataSource3 := make(chan Task, 10)

	go func() {
		defer close(dataSource1)
		for i := 1; i <= 5; i++ {
			dataSource1 <- Task{ID: i, Data: fmt.Sprintf("Data from Source 1-%d", i), ProcessTime: time.Duration(rand.Intn(500)) * time.Millisecond}
		}
	}()

	go func() {
		defer close(dataSource2)
		for i := 1; i <= 3; i++ {
			dataSource2 <- Task{ID: i, Data: fmt.Sprintf("Data from Source 2-%d", i), ProcessTime: time.Duration(rand.Intn(500)) * time.Millisecond}
		}
	}()

	go func() {
		defer close(dataSource3)
		for i := 1; i <= 4; i++ {
			dataSource3 <- Task{ID: i, Data: fmt.Sprintf("Data from Source 3-%d", i), ProcessTime: time.Duration(rand.Intn(500)) * time.Millisecond}
		}
	}()

	// Define the number of workers
	numWorkers := 3

	// Manage tasks concurrently from multiple data sources
	manageTasks(numWorkers, dataSource1, dataSource2, dataSource3)

	fmt.Println("All tasks processed concurrently.")
}
