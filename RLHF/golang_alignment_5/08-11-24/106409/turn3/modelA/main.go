package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Counter represents a shared counter that can be incremented and decremented concurrently.
type Counter struct {
	value int
	mutex sync.RWMutex
}

// NewCounter creates a new Counter instance.
func NewCounter() *Counter {
	return &Counter{}
}

// Incr increments the counter safely.
func (c *Counter) Incr() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

// Decr decrements the counter safely.
func (c *Counter) Decr() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value--
}

// Read reads the current value of the counter safely.
func (c *Counter) Read() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.value
}

// simulateConcurrentAccess simulates multiple concurrent read and write operations on the shared counter.
func simulateConcurrentAccess(wg *sync.WaitGroup, counter *Counter) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		if rand.Intn(2) == 0 {
			// Simulate read operation
			currentValue := counter.Read()
			fmt.Printf("Read: %d\n", currentValue)
		} else {
			// Simulate write operation (increment or decrement)
			if rand.Intn(2) == 0 {
				counter.Incr()
				fmt.Println("Incr")
			} else {
				counter.Decr()
				fmt.Println("Decr")
			}
		}
	}
}

func main() {
	counter := NewCounter()
	var wg sync.WaitGroup
	numConcurrentAccesses := 5
	wg.Add(numConcurrentAccesses)

	for i := 0; i < numConcurrentAccesses; i++ {
		go simulateConcurrentAccess(&wg, counter)
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter.Read())
}
