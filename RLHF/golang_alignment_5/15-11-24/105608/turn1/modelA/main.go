package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	numWorkers  = 10
	numTasks    = 1000
	taskTimeout = time.Millisecond * 100 // Random task duration
)

var (
	taskQueue         = make(chan int, numTasks)
	done              int32
	workGroup         sync.WaitGroup
	wgComputeDuration int64 = 0
)

// worker is a goroutine that processes tasks from the queue.
func worker() {
	for {
		task, ok := <-taskQueue
		if !ok {
			break
		}

		// Simulate work being done
		computeDuration := rand.Intn(int(taskTimeout.Nanoseconds()))
		time.Sleep(time.Duration(computeDuration) * time.Nanosecond)
		wgComputeDuration += int64(computeDuration)

		// Signal task completion
		atomic.AddInt32(&done, 1)
	}
}

func main() {
	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// Start worker goroutines
	workGroup.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker()
	}

	// Enqueue tasks
	for i := 0; i < numTasks; i++ {
		taskQueue <- i
	}

	// Close the task queue to signal workers to exit
	close(taskQueue)

	// Wait for all workers to finish
	workGroup.Wait()

	// Calculate average compute time and report stats
	average := float64(wgComputeDuration) / float64(done)
	fmt.Printf("Average compute duration: %.3fμs\n", average/1000.0)
	fmt.Printf("Num Tasks: %d\n", numTasks)
	fmt.Printf("Num Workers: %d\n", numWorkers)
}
