package main

import (
	"fmt"
	"sync"
)

// DistributedCounter represents a distributed counter.
type DistributedCounter struct {
	replicas []*Replica
	mu       sync.Mutex
}

// Replica represents a single replica of the distributed counter.
type Replica struct {
	value int
}

// NewDistributedCounter creates a new distributed counter with the specified number of replicas.
func NewDistributedCounter(numReplicas int) *DistributedCounter {
	dc := &DistributedCounter{}
	for i := 0; i < numReplicas; i++ {
		dc.replicas = append(dc.replicas, &Replica{})
	}
	return dc
}

// Increment increments the counter by 1 and returns the new value.
func (dc *DistributedCounter) Increment() int {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	// Increment the counter on all replicas
	for _, replica := range dc.replicas {
		replica.value++
	}

	// Return the updated value from the first replica
	return dc.replicas[0].value
}

// GetValue retrieves the current value of the counter.
func (dc *DistributedCounter) GetValue() int {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	// Read the value from the first replica
	return dc.replicas[0].value
}

func main() {
	// Create a distributed counter with 3 replicas
	counter := NewDistributedCounter(3)
	fmt.Println("Counter :", counter)
}
