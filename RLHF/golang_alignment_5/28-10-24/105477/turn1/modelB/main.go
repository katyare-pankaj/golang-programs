package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// SynchronizedData defines the structure of the data being synchronized
type SynchronizedData struct {
	Timestamp int64
	Data      []int
}

// Synchronizer manages synchronization of the data
type Synchronizer struct {
	mu        sync.Mutex
	data      SynchronizedData
	listeners []chan SynchronizedData
}

// Init initializes the synchronizer with initial data and an empty slice of listeners
func (s *Synchronizer) Init(timestamp int64, data []int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = SynchronizedData{Timestamp: timestamp, Data: data}
	s.listeners = make([]chan SynchronizedData, 0)
}

// Subscribe adds a listener channel to the synchronizer's list of listeners
func (s *Synchronizer) Subscribe(ch chan SynchronizedData) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.listeners = append(s.listeners, ch)
}

// UpdateData updates the synchronized data and notifies all listeners
func (s *Synchronizer) UpdateData(timestamp int64, data []int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = SynchronizedData{Timestamp: timestamp, Data: data}
	for _, ch := range s.listeners {
		ch <- s.data
	}
}

// LatestData returns the latest synchronized data
func (s *Synchronizer) LatestData() SynchronizedData {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.data
}

// *** Testing functions ***

// createRandomData generates random data for testing
func createRandomData(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(100)
	}
	return data
}

// synchronizationBenchmark tests the performance of synchronization using a loop
func synchronizationBenchmark(b *testing.B, numUpdates int, listenerChan chan SynchronizedData) {
	syncer := new(Synchronizer)
	syncer.Init(time.Now().UnixNano(), createRandomData(10))
	syncer.Subscribe(listenerChan)

	b.ResetTimer()
	for i := 0; i < numUpdates; i++ {
		data := createRandomData(10)
		syncer.UpdateData(time.Now().UnixNano(), data)
	}
}

func TestSynchronization(t *testing.T) {
	syncer := new(Synchronizer)
	syncer.Init(time.Now().UnixNano(), createRandomData(10))

	// Create a listener channel
	listenerChan := make(chan SynchronizedData)
	syncer.Subscribe(listenerChan)

	// Test initial data synchronization
	expectedData := syncer.LatestData()
	select {
	case actualData := <-listenerChan:
		if !equalSynchronizedData(expectedData, actualData) {
			t.Errorf("Initial data synchronization failed. Expected: %v, Got: %v", expectedData, actualData)
		}
	case <-time.After(time.Second):
		t.Errorf("Initial data synchronization timed out")
	}

	// Perform synchronization benchmark
	numUpdates := 10000
	b := &testing.B{}
	synchronizationBenchmark(b, numUpdates, listenerChan)

	fmt.Printf("Synchronization benchmark completed with %d updates\n", numUpdates)
}
