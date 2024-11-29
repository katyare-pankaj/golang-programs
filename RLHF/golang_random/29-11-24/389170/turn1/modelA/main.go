package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Worker struct {
	id       int
	work     chan<- string
	done     chan bool
	wg       *sync.WaitGroup
	shutdown chan struct{}
}

func NewWorker(id int, work chan<- string, wg *sync.WaitGroup) *Worker {
	return &Worker{
		id:       id,
		work:     work,
		done:     make(chan bool),
		wg:       wg,
		shutdown: make(chan struct{}),
	}
}

func (w *Worker) Start() {
	defer w.wg.Done()

	log.Printf("Worker %d started\n", w.id)
	for {
		select {
		case task := <-w.work:
			log.Printf("Worker %d processing task: %s\n", w.id, task)
			time.Sleep(time.Second) // Simulate work
			log.Printf("Worker %d completed task: %s\n", w.id, task)
		case <-w.shutdown:
			log.Printf("Worker %d shutting down\n", w.id)
			close(w.done)
			return
		}
	}
}

func main() {
	const (
		numWorkers = 3
		numTasks   = 10
	)

	// Create a channel for work
	work := make(chan string, numTasks)

	// Create and start workers
	wg := &sync.WaitGroup{}
	workers := make([]*Worker, numWorkers)
	for i := range workers {
		workers[i] = NewWorker(i, work, wg)
		go workers[i].Start()
	}

	// Submit tasks
	for i := 1; i <= numTasks; i++ {
		work <- fmt.Sprintf("Task %d", i)
	}

	// Close the work channel when all tasks are submitted
	close(work)

	// Wait for all workers to finish
	wg.Wait()

	// Shutdown workers
	for _, worker := range workers {
		worker.shutdown <- struct{}{}
		<-worker.done
	}

	log.Println("All workers completed")
}
