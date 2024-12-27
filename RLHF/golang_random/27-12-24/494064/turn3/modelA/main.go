package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type WorkerPool struct {
	workers    []chan<- int
	available  int32
	maxWorkers int
	wg         *sync.WaitGroup
	closeChan  chan struct{}
}

// NewWorkerPool initializes a new worker pool with a fixed number of workers.
func NewWorkerPool(maxWorkers int) *WorkerPool {
	pool := &WorkerPool{
		maxWorkers: maxWorkers,
		workers:    make([]chan<- int, maxWorkers),
		available:  int32(maxWorkers),
		wg:         &sync.WaitGroup{},
		closeChan:  make(chan struct{}),
	}

	for i := 0; i < maxWorkers; i++ {
		workerChan := make(chan int)
		pool.workers[i] = workerChan
		pool.wg.Add(1)

		go func(workerID int) {
			for {
				select {
				case task, ok := <-workerChan:
					if !ok {
						return
					}
					fmt.Printf("Worker %d started task %d\n", workerID, task)
					time.Sleep(time.Second * 2) // Simulate work
					fmt.Printf("Worker %d finished task %d\n", workerID, task)
				case <-pool.closeChan:
					fmt.Printf("Worker %d is shutting down\n", workerID)
					return
				}
			}
		}(i)
	}

	return pool
}

// Submit adds a task to the worker pool.
func (pool *WorkerPool) Submit(task int) bool {
	if pool.available <= 0 {
		return false // No available workers
	}

	for {
		workerID := int(atomic.AddInt32(&pool.available, -1))
		select {
		case pool.workers[workerID] <- task:
			return true
		default:
			atomic.AddInt32(&pool.available, 1)
			time.Sleep(10 * time.Millisecond) // Back off slightly before retrying
		}
	}
}

// Close stops the worker pool.
func (pool *WorkerPool) Close() {
	close(pool.closeChan)
	pool.wg.Wait()
}

func main() {
	pool := NewWorkerPool(5)
	defer pool.Close()

	for i := 0; i < 20; i++ {
		if pool.Submit(i + 1) {
			fmt.Printf("Submitted task %d\n", i+1)
		} else {
			fmt.Printf("Failed to submit task %d (no available workers)\n", i+1)
		}
	}

	fmt.Println("Waiting for tasks to complete...")
}
