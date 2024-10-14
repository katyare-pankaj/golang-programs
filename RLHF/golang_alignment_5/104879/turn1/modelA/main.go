package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers    = 10
	taskQueueSize = 100
)

// Simulate API request
type apiRequest struct {
	id int
}

// Simulate API response
type apiResponse struct {
	request apiRequest
	data    string
}

// Worker represents a goroutine that processes API requests
type worker struct {
	id  int
	in  chan apiRequest
	out chan apiResponse
	wg  sync.WaitGroup
}

// NewWorker creates a new worker
func newWorker(id int, in chan apiRequest, out chan apiResponse, wg *sync.WaitGroup) *worker {
	w := &worker{
		id:  id,
		in:  in,
		out: out,
		wg:  *wg,
	}
	w.wg.Add(1)
	go w.work()
	return w
}

// work is the main loop for a worker goroutine
func (w *worker) work() {
	defer w.wg.Done()
	for request := range w.in {
		// Simulate processing time
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		response := apiResponse{request: request, data: fmt.Sprintf("Response for request %d", request.id)}
		w.out <- response
	}
}

func main() {
	// Create task queue and response channel
	taskQueue := make(chan apiRequest, taskQueueSize)
	responseChan := make(chan apiResponse)

	var wg sync.WaitGroup

	// Create workers
	for i := 0; i < numWorkers; i++ {
		newWorker(i, taskQueue, responseChan, &wg)
	}

	// Generate and send API requests
	go func() {
		for i := 0; i < 100; i++ { // Generate 100 requests
			taskQueue <- apiRequest{id: i}
		}
		close(taskQueue) // Close the task queue to signal workers to exit
	}()

	// Collect and print responses
	go func() {
		for response := range responseChan {
			fmt.Printf("Worker %d processed request %d: %s\n", response.request.id%numWorkers, response.request.id, response.data)
		}
		wg.Done()
	}()

	wg.Wait() // Wait for all workers to finish
	fmt.Println("All tasks completed.")
}
