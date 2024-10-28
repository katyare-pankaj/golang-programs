package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

const (
	numVehicles   = 100
	numIterations = 10000
)

type VectorClock struct {
	mu      sync.Mutex
	clocks  map[int]int
	vectors map[int]map[int]int
}

func NewVectorClock() *VectorClock {
	return &VectorClock{
		clocks:  make(map[int]int),
		vectors: make(map[int]map[int]int),
	}
}

func (vc *VectorClock) Synchronize(vehicleID int, currentTime int) {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	// Update the local clock
	vc.clocks[vehicleID] = max(vc.clocks[vehicleID], currentTime)

	// Merge received vector clocks
	for senderID, senderVector := range vc.vectors {
		vc.clocks[senderID] = max(vc.clocks[senderID], senderVector[vehicleID])
	}

	// Update vectors
	vc.vectors[vehicleID] = make(map[int]int)
	for senderID, senderTime := range vc.clocks {
		vc.vectors[vehicleID][senderID] = senderTime
	}
}

func (vc *VectorClock) GetSynchronizedTime(vehicleID int) int {
	vc.mu.Lock()
	defer vc.mu.Unlock()
	return vc.clocks[vehicleID]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func TestVectorClock(t *testing.T) {
	vc := NewVectorClock()

	// Simulate synchronization between vehicles
	for vehicleID := 0; vehicleID < numVehicles; vehicleID++ {
		for i := 0; i < numIterations; i++ {
			currentTime := rand.Intn(10000)
			vc.Synchronize(vehicleID, currentTime)
		}
	}

	// Ensure all vehicles have the same synchronized time
	synchronizedTime := vc.GetSynchronizedTime(0)
	for vehicleID := 1; vehicleID < numVehicles; vehicleID++ {
		if vc.GetSynchronizedTime(vehicleID) != synchronizedTime {
			t.Errorf("Vehicles %d and %d have different synchronized times", 0, vehicleID)
		}
	}
}

func main() {
	// Run performance test for the VectorClock
	fmt.Println("Testing VectorClock:")
	TestVectorClock(testing.T{})
}
