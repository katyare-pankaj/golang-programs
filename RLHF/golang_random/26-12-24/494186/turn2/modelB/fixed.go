package main

import (
	"errors"
	"fmt"
	"sync"
)

// Counter struct with a mutex to protect the value
type Counter struct {
	value int
	sync.Mutex
}

// Increment method to increment the counter safely
func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

// Decrement method to decrement the counter safely
func (c *Counter) Decrement() {
	c.Lock()
	defer c.Unlock()
	c.value--
}

func worker(id int, c *Counter, done chan bool, errChan chan error) {
	defer close(done)
	defer close(errChan)
	for i := 0; i < 1000; i++ {
		c.Increment()
		c.Decrement()
	}
	// Simulate a random error
	if id == 3 {
		errChan <- errors.New("worker 3 failed")
		return
	}
}

func main() {
	numWorkers := 5
	c := &Counter{}
	done := make(chan bool)
	errChan := make(chan error)
	for i := 1; i <= numWorkers; i++ {
		go worker(i, c, done, errChan)
	}
	// Wait for all workers to complete or report an error
	for i := 1; i <= numWorkers; i++ {
		select {
		case <-done:
			fmt.Println("Worker", i, "done")
		case err := <-errChan:
			fmt.Println("Error:", err)
		}
	}
	// Read the final counter value
	c.Lock()
	finalValue := c.value
	c.Unlock()
	fmt.Println("Final Counter Value:", finalValue)
}
