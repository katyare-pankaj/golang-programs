package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers    = 4    // Number of worker goroutines
	bufferSize    = 1000 // Buffer size for worker channels
	maxLoadFactor = 0.8  // Maximum load factor before spawning a new worker
)

type Worker struct {
	id          int
	taskChannel chan int
	wg          *sync.WaitGroup
	quit        chan struct{}
}

func newWorker(id int, wg *sync.WaitGroup) *Worker {
	return &Worker{
		id:          id,
		taskChannel: make(chan int, bufferSize),
		wg:          wg,
		quit:        make(chan struct{}),
	}
}

func (w *Worker) start() {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		for {
			select {
			case task, ok := <-w.taskChannel:
				if !ok {
					return // Worker channel is closed, exit
				}
				// Simulate processing data
				time.Sleep(time.Duration(task) * time.Millisecond)
				fmt.Printf("Worker %d processed task %d\n", w.id, task)
			case <-w.quit:
				return // Worker is quitting
			}
		}
	}()
}

func (w *Worker) stop() {
	close(w.taskChannel) // Close the worker channel to signal termination
	<-w.quit             // Wait for the worker to finish
}

type LoadBalancer struct {
	workers     []*Worker
	nextWorker  int
	loadFactor  float32
	wg          *sync.WaitGroup
	taskChannel chan int
}

func newLoadBalancer(wg *sync.WaitGroup) *LoadBalancer {
	lb := &LoadBalancer{
		workers:     make([]*Worker, 0, numWorkers),
		nextWorker:  0,
		loadFactor:  0,
		wg:          wg,
		taskChannel: make(chan int, bufferSize),
	}

	// Start the initial set of workers
	for i := 0; i < numWorkers; i++ {
		w := newWorker(i, lb.wg)
		w.start()
		lb.workers = append(lb.workers, w)
	}

	return lb
}

func (lb *LoadBalancer) start() {
	lb.wg.Add(1)
	go func() {
		defer lb.wg.Done()
		for {
			task, ok := <-lb.taskChannel
			if !ok {
				return // Load balancer channel is closed, exit
			}

			// Distribute tasks to workers using round-robin
			worker := lb.workers[lb.nextWorker]
			lb.nextWorker = (lb.nextWorker + 1) % len(lb.workers)

			// Check if worker is overloaded and spawn a new one if needed
			lb.checkLoad()

			select {
			case worker.taskChannel <- task:
			default:
				// Worker channel is full, try again
				lb.taskChannel <- task
			}
		}
	}()
}

func (lb *LoadBalancer) stop() {
	close(lb.taskChannel) // Close the load balancer channel to signal termination
	for _, worker := range lb.workers {
		worker.stop() // Stop each worker
	}
}

func (lb *LoadBalancer) checkLoad() {
	// Calculate the average load factor across all workers
	totalTasks := 0
	for _, worker := range lb.workers {
		totalTasks += len(worker.taskChannel)
	}
	lb.loadFactor = float32(totalTasks) / float32(numWorkers*bufferSize)

	// If load factor is above the threshold, spawn a new worker
	if lb.loadFactor > maxLoadFactor {
		w := newWorker(len(lb.workers), lb.wg)
		w.start()
		lb.workers = append(lb.workers, w)
	}
}

func main() {
	var wg sync.WaitGroup
	lb := newLoadBalancer(&wg)
	lb.start()

	// Generate random tasks for simulation
	go func() {
		for i := 0; i < 100; i++ {
			task := rand.Intn(50) + 1 // Generate random task duration between 1 and 500 ms
			lb.taskChannel <- task
		}
		close(lb.taskChannel)
	}()

	wg.Wait() // Wait for all workers and the load balancer to finish
}
