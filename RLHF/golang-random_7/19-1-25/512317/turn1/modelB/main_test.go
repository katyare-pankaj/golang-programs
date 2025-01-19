package main

import (
	"sync"
	"testing"
)

// Mutable state implementation
type MutableCounter struct {
	count int
	mu    sync.Mutex
}

func (c *MutableCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *MutableCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// Immutable state implementation
type ImmutableCounter struct {
	count int
}

func (c *ImmutableCounter) Increment() *ImmutableCounter {
	return &ImmutableCounter{count: c.count + 1}
}

func (c *ImmutableCounter) GetCount() int {
	return c.count
}

// Benchmarking functions
func BenchmarkMutableCallbacks(b *testing.B) {
	counter := &MutableCounter{}
	for n := 0; n < b.N; n++ {
		counter.Increment()
	}
}

func BenchmarkImmutableCallbacks(b *testing.B) {
	counter := &ImmutableCounter{}
	for n := 0; n < b.N; n++ {
		counter = counter.Increment()
	}
}

// Test function to check data integrity
func TestCallbackIntegrity(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	// Mutable state test
	mutableCounter := &MutableCounter{}
	go func() {
		defer wg.Done()
		for n := 0; n < 10000; n++ {
			mutableCounter.Increment()
		}
	}()

	go func() {
		defer wg.Done()
		for n := 0; n < 10000; n++ {
			mutableCounter.Increment()
		}
	}()

	wg.Wait()
	expectedCount := 20000
	actualCount := mutableCounter.GetCount()
	if actualCount != expectedCount {
		t.Errorf("Mutable counter integrity failed: Expected %d, got %d", expectedCount, actualCount)
	}

	// Immutable state test
	immutableCounter := &ImmutableCounter{}
	for n := 0; n < 10000; n++ {
		immutableCounter = immutableCounter.Increment()
	}

	for n := 0; n < 10000; n++ {
		immutableCounter = immutableCounter.Increment()
	}

	expectedCount = 20000
	actualCount = immutableCounter.GetCount()
	if actualCount != expectedCount {
		t.Errorf("Immutable counter integrity failed: Expected %d, got %d", expectedCount, actualCount)
	}
}
