package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	tokens chan struct{}
}

// NewSemaphore creates a new Semaphore with a specified limit
func NewSemaphore(limit int) *Semaphore {
	return &Semaphore{
		tokens: make(chan struct{}, limit), // Buffered channel
	}
}

// Acquire a token from the semaphore
func (s *Semaphore) Acquire() {
	s.tokens <- struct{}{} // Block if the channel is full
}

// Release a token back to the semaphore
func (s *Semaphore) Release() {
	<-s.tokens // Unblock if there are tokens available
}

func worker(id int, sem *Semaphore, wg *sync.WaitGroup) {
	defer wg.Done()
	sem.Acquire() // Acquire a token

	// Simulate work
	fmt.Printf("Worker %d is starting work\n", id)
	time.Sleep(2 * time.Second) // Simulating work
	fmt.Printf("Worker %d is done\n", id)

	sem.Release() // Release the token
}

func main() {
	const numWorkers = 5
	const semLimit = 2 // Limit of concurrent workers

	sem := NewSemaphore(semLimit)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, sem, &wg)
	}

	wg.Wait() // Wait for all workers to finish
	fmt.Println("All workers have completed their tasks")
}
