package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Semaphore represents a semaphore implemented using a channel.
type Semaphore struct {
	ch chan struct{}
}

// NewSemaphore creates a new semaphore with the specified value.
func NewSemaphore(value int) *Semaphore {
	return &Semaphore{ch: make(chan struct{}, value)}
}

// Acquire acquires the semaphore.
func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

// Release releases the semaphore.
func (s *Semaphore) Release() {
	<-s.ch
}

func worker(sem *Semaphore, wg *sync.WaitGroup, id int, task func()) {
	defer wg.Done()
	for {
		// Acquire the semaphore
		sem.Acquire()
		defer sem.Release()

		// Execute the task
		task()

		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	sem := NewSemaphore(5) // Limit the number of concurrent workers to 5

	// Define a task to simulate work
	task := func() {
		fmt.Printf("Working...\n")
	}

	// Start 10 worker goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(sem, &wg, i, task)
	}

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers finished.")
}
