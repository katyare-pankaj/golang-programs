package main

import (
	"fmt"
	"sync"
	"time"
)

// ResourceManager handles a shared resource with a dedicated Goroutine
type ResourceManager struct {
	requests chan struct{}
	counter  int32
	mu       sync.Mutex
}

// NewResourceManager initializes the ResourceManager
func NewResourceManager(bufferSize int) *ResourceManager {
	rm := &ResourceManager{
		requests: make(chan struct{}, bufferSize),
	}
	go rm.processRequests()
	return rm
}

// RequestResource simulates a constituting request for the resource
func (rm *ResourceManager) RequestResource(id int) {
	rm.mu.Lock()
	rm.counter++
	rm.mu.Unlock()

	rm.requests <- struct{}{}
	// Simulate processing the resource
	time.Sleep(100 * time.Millisecond)
	<-rm.requests
}

// processRequests handles the requests in a single Goroutine
func (rm *ResourceManager) processRequests() {
	for {
		<-rm.requests
		// Simulate doing work with the shared resource
		time.Sleep(50 * time.Millisecond) // Simulated work time
	}
}

// Metrics gathers the channel utilization metrics
func (rm *ResourceManager) Metrics() (int, int32) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	return len(rm.requests), rm.counter
}

func main() {
	const totalRequests = 100
	const numWorkers = 10
	rm := NewResourceManager(1)

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < totalRequests/numWorkers; j++ {
				rm.RequestResource(id)
			}
		}(i)
	}

	wg.Wait()

	// Display metrics
	queued, total := rm.Metrics()
	fmt.Printf("Requests in Queue: %d, Total Accesses: %d\n", queued, total)
}
