package main

import (
	"fmt"
	"sync"
	"time"
)

// Synchronizer provides a mechanism to synchronize multiple goroutines
type Synchronizer interface {
	Synchronize(func())
}

// BasicSynchronizer uses a mutex to ensure only one goroutine can execute the critical section at a time
type BasicSynchronizer struct {
	mu sync.Mutex
}

func (s *BasicSynchronizer) Synchronize(f func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	f()
}

// ChannelSynchronizer uses a channel to allow only one goroutine to proceed at a time
type ChannelSynchronizer struct {
	ch chan struct{}
}

func (s *ChannelSynchronizer) Synchronize(f func()) {
	s.ch <- struct{}{}
	defer func() { <-s.ch }()
	f()
}

// TestSynchronizer tests the performance of a Synchronizer
func TestSynchronizer(s Synchronizer, numGoroutines, numIterations int) {
	var wg sync.WaitGroup
	startTime := time.Now()

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numIterations; j++ {
				s.Synchronize(func() {
					// Critical section code goes here
					_ = j
				})
			}
		}()
	}

	wg.Wait()
	duration := time.Since(startTime)
	fmt.Printf("Test completed in %s\n", duration)
}

func main() {
	const numGoroutines = 100
	const numIterations = 100000

	fmt.Println("Testing BasicSynchronizer:")
	basicSynchronizer := &BasicSynchronizer{}
	TestSynchronizer(basicSynchronizer, numGoroutines, numIterations)

	fmt.Println("\nTesting ChannelSynchronizer:")
	channelSynchronizer := &ChannelSynchronizer{ch: make(chan struct{}, 1)}
	TestSynchronizer(channelSynchronizer, numGoroutines, numIterations)
}
