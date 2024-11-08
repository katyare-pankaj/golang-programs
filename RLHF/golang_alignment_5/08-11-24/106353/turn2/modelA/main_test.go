package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

// ScalableInventorySystem represents a scalable inventory management system
type ScalableInventorySystem struct {
	mutex     sync.RWMutex
	inventory map[string]int
}

// NewScalableInventorySystem creates a new ScalableInventorySystem
func NewScalableInventorySystem() *ScalableInventorySystem {
	return &ScalableInventorySystem{
		inventory: make(map[string]int),
	}
}

// UpdateItem updates the quantity of an item in the inventory
func (s *ScalableInventorySystem) UpdateItem(item string, quantity int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.inventory[item] += quantity
}

// GetItemQuantity retrieves the quantity of an item from the inventory
func (s *ScalableInventorySystem) GetItemQuantity(item string) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.inventory[item]
}

// TestScalableInventorySystem tests the ScalableInventorySystem for correctness and scalability
func TestScalableInventorySystem(t *testing.T) {
	const numItems = 1000
	const numUpdates = 10000
	const numGoroutines = 10

	sys := NewScalableInventorySystem()

	// Initialize inventory
	for i := 0; i < numItems; i++ {
		sys.UpdateItem(fmt.Sprintf("item%d", i), 10)
	}

	// Perform concurrent updates
	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numUpdates; j++ {
				itemIndex := rand.Intn(numItems)
				quantity := rand.Intn(20) - 10
				sys.UpdateItem(fmt.Sprintf("item%d", itemIndex), quantity)
			}
		}()
	}

	wg.Wait()

	// Validate final quantities
	for i := 0; i < numItems; i++ {
		expectedQuantity := 10
		for j := 0; j < numUpdates; j++ {
			quantity := rand.Intn(20) - 10
			expectedQuantity += quantity
		}
		if actualQuantity := sys.GetItemQuantity(fmt.Sprintf("item%d", i)); actualQuantity != expectedQuantity {
			t.Errorf("Item %d quantity mismatch: expected %d, got %d", i, expectedQuantity, actualQuantity)
		}
	}
}

func main() {
	// Run tests
	testing.Init()
	testing.MainStart(nil, nil)
	defer testing.MainEnd()

	tests := []testing.InternalTest{{Name: "TestScalableInventorySystem", F: TestScalableInventorySystem}}
	for _, test := range tests {
		if test.F != nil {
			fmt.Printf("Running test: %s\n", test.Name)
			test.F(nil, nil)
		}
	}
}
