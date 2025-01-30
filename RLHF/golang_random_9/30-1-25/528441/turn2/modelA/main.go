package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	maxWorkers         = 10
	minWorkers         = 2
	taskQueueSize      = 100
	scaleUpThreshold   = 75
	scaleDownThreshold = 25
	taskSimulationTime = 500 // in milliseconds
)

type Task struct {
	id int
}

func worker(id int, taskQueue <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskQueue {
		fmt.Printf("Worker %d processing task %d\n", id, task.id)
		time.Sleep(time.Duration(rand.Intn(taskSimulationTime)) * time.Millisecond) // Simulate work
	}
}

func loadBalancer(taskQueue chan Task, wg *sync.WaitGroup) {
	activeWorkers := minWorkers
	workerWG := &sync.WaitGroup{}

	// Start initial worker pool
	for i := 0; i < activeWorkers; i++ {
		workerWG.Add(1)
		go worker(i, taskQueue, workerWG)
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		queueLength := len(taskQueue)
		fmt.Printf("Task Queue Length: %d\n", queueLength)

		if queueLength > scaleUpThreshold && activeWorkers < maxWorkers {
			// Scale up
			fmt.Println("Scaling up workers...")
			activeWorkers++
			workerWG.Add(1)
			go worker(activeWorkers-1, taskQueue, workerWG)
		} else if queueLength < scaleDownThreshold && activeWorkers > minWorkers {
			// Scale down
			fmt.Println("Scaling down workers...")
			activeWorkers--
			close(taskQueue)
			taskQueue = make(chan Task, taskQueueSize)
			workerWG.Wait()

			// Restart remaining workers
			for i := 0; i < activeWorkers; i++ {
				workerWG.Add(1)
				go worker(i, taskQueue, workerWG)
			}
		}
	}
}

func main() {
	taskQueue := make(chan Task, taskQueueSize)
	var wg sync.WaitGroup

	// Start load balancer
	go loadBalancer(taskQueue, &wg)

	// Generate tasks
	for i := 0; i < 100; i++ {
		wg.Add(1)
		taskQueue <- Task{id: i}
	}

	// Close the task queue when done adding tasks
	close(taskQueue)

	// Wait for all tasks to be processed
	wg.Wait()
	fmt.Println("All tasks processed")
}
