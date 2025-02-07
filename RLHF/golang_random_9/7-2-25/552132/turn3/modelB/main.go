package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWriters  = 3   // Number of goroutines writing to the slice
	numReaders  = 10  // Number of goroutines reading from the slice
	numElements = 100 // Total number of elements in the slice
	readSpeed   = 10  // Duration between reads (in milliseconds)
	writeSpeed  = 50  // Duration between writes (in milliseconds)
)

// SharedSlice represents a slice with an RWMutex for synchronized access.
type SharedSlice struct {
	data  []int
	mutex sync.RWMutex
}

// NewSharedSlice initializes a SharedSlice with random data.
func NewSharedSlice(size int) *SharedSlice {
	s := &SharedSlice{
		data: make([]int, size),
	}
	for i := range s.data {
		s.data[i] = rand.Intn(1000)
	}
	return s
}

// Write generates a random number and adds it to the slice safely.
func (s *SharedSlice) Write() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data = append(s.data, rand.Intn(1000))
}

// Read retrieves a random element from the slice safely.
func (s *SharedSlice) Read() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if len(s.data) == 0 {
		return -1
	}
	return s.data[rand.Intn(len(s.data))]
}

func writer(id int, sharedSlice *SharedSlice, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		sharedSlice.Write()
		time.Sleep(time.Millisecond * writeSpeed)
	}
}

func reader(id int, sharedSlice *SharedSlice, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		value := sharedSlice.Read()
		if value != -1 {
			fmt.Printf("Reader %d read: %d\n", id, value)
		}
		time.Sleep(time.Millisecond * readSpeed)
	}
}

func main() {
	sharedSlice := NewSharedSlice(numElements)
	var wg sync.WaitGroup

	// Launch writers
	for i := 0; i < numWriters; i++ {
		wg.Add(1)
		go writer(i, sharedSlice, &wg)
	}

	// Launch readers
	for i := 0; i < numReaders; i++ {
		wg.Add(1)
		go reader(i, sharedSlice, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
