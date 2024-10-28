package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

const (
	numVehicles   = 100
	numIterations = 10000
)

// SynchronizationMethod represents the interface for different synchronization methods
type SynchronizationMethod interface {
	Synchronize(vehicleID int, currentTime int) int
}

// BasicSynchronization implements a simple synchronization method
type BasicSynchronization struct {
	mu                   sync.Mutex
	lastSynchronizedTime int
}

func (s *BasicSynchronization) Synchronize(vehicleID int, currentTime int) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	if currentTime > s.lastSynchronizedTime {
		s.lastSynchronizedTime = currentTime
	}
	return s.lastSynchronizedTime
}

// TestSynchronizationMethod performs performance testing for a given SynchronizationMethod
func TestSynchronizationMethod(t *testing.T, method SynchronizationMethod) {
	var wg sync.WaitGroup
	startTime := time.Now()

	wg.Add(numVehicles)
	for vehicleID := 0; vehicleID < numVehicles; vehicleID++ {
		go func(vehicleID int) {
			defer wg.Done()
			for i := 0; i < numIterations; i++ {
				currentTime := rand.Intn(10000)
				method.Synchronize(vehicleID, currentTime)
			}
		}(vehicleID)
	}

	wg.Wait()
	duration := time.Since(startTime)

	fmt.Printf("Test completed in %s\n", duration)
}

func main() {
	// Create instances of the synchronization method to be tested
	basicSync := &BasicSynchronization{}

	// Run performance test for the BasicSynchronization method
	fmt.Println("Testing BasicSynchronization:")
	TestSynchronizationMethod(testing.T{}, basicSync)
}
