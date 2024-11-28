package main

import (
	"fmt"
	"sync"
	"time"
)

const workerPoolSize = 10

type workerPool struct {
	workers  chan chan<- int
	tasks    chan int
	shutdown chan struct{}
	wg       *sync.WaitGroup
}

func newWorkerPool() *workerPool {
	wp := &workerPool{
		workers:  make(chan chan<- int, workerPoolSize),
		tasks:    make(chan int),
		shutdown: make(chan struct{}),
		wg:       &sync.WaitGroup{},
	}

	// Start worker goroutines
	for i := 0; i < workerPoolSize; i++ {
		workerChan := make(chan int)
		wp.workers <- workerChan

		wp.wg.Add(1)
		go wp.worker(workerChan)
	}

	return wp
}

func (wp *workerPool) worker(workerChan chan<- int) {
	defer wp.wg.Done()

	for {
		select {
		case task, ok := <-workerChan:
			if !ok {
				return
			}
			fmt.Printf("Worker %d processing task %d\n", workerChan, task)
			time.Sleep(time.Duration(task) * time.Millisecond)

		case <-wp.shutdown:
			return
		}
	}
}

func (wp *workerPool) submitTask(task int) {
	select {
	case workerChan := <-wp.workers:
		workerChan <- task
	case <-wp.shutdown:
		// Ignore task if the pool is shutting down
	}
}

func (wp *workerPool) shutdownPool() {
	close(wp.workers)
	close(wp.tasks)
	wp.wg.Wait()
	close(wp.shutdown)
}

func main() {
	wp := newWorkerPool()

	// Submit tasks
	for i := 1; i <= 100; i++ {
		wp.submitTask(i)
	}

	// Shut down the pool
	wp.shutdownPool()
}
