package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numWorkers   = 10
	numResources = 3
)

// ResourcePool manages a set of resources.
type ResourcePool struct {
	resources chan string
	mu        sync.Mutex
}

// NewResourcePool creates a new resource pool with a given capacity.
func NewResourcePool(capacity int) *ResourcePool {
	return &ResourcePool{
		resources: make(chan string, capacity),
	}
}

// GetResource acquires a resource from the pool.
func (p *ResourcePool) GetResource() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	return <-p.resources
}

// ReleaseResource releases a resource back to the pool.
func (p *ResourcePool) ReleaseResource(r string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.resources <- r
}

// worker performs some work with a resource.
func worker(id int, pool *ResourcePool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		fmt.Printf("Worker %d: Acquiring resource\n", id)
		// Acquire a resource
		res := pool.GetResource()
		defer pool.ReleaseResource(res)

		fmt.Printf("Worker %d: Using resource %s\n", id, res)
		// Simulate work using the resource
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	pool := NewResourcePool(numResources)

	// Initialize the resource pool with available resources
	for i := 1; i <= numResources; i++ {
		pool.resources <- fmt.Sprintf("Resource%d", i)
	}

	wg.Add(numWorkers)
	// Start workers that will acquire and use resources from the pool
	for i := 1; i <= numWorkers; i++ {
		go worker(i, pool, &wg)
	}

	fmt.Println("Workers started...")
	wg.Wait()
	fmt.Println("Workers completed.")
}
