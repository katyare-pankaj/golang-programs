package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeSlice encapsulates a slice and uses RWMutex for concurrent access.
type SafeSlice struct {
	sync.RWMutex
	data []int
}

// Add safely appends a value to the slice.
func (s *SafeSlice) Add(value int) {
	s.Lock()
	defer s.Unlock()
	s.data = append(s.data, value)
}

// Get safely retrieves a copy of the slice's data.
func (s *SafeSlice) Get() []int {
	s.RLock()
	defer s.RUnlock()
	// Return a copy of the slice to avoid race conditions
	copyData := make([]int, len(s.data))
	copy(copyData, s.data)
	return copyData
}

func writer(id int, wg *sync.WaitGroup, sharedSlice *SafeSlice) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		sharedSlice.Add(id*10 + i)
		fmt.Printf("Writer %d added value %d\n", id, id*10+i)
		time.Sleep(time.Millisecond * 100) // Simulate work
	}
}

func reader(id int, wg *sync.WaitGroup, sharedSlice *SafeSlice) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		data := sharedSlice.Get()
		fmt.Printf("Reader %d read data: %v\n", id, data)
		time.Sleep(time.Millisecond * 150) // Simulate work
	}
}

func main() {
	var wg sync.WaitGroup
	sharedSlice := &SafeSlice{data: []int{}}

	// Launch writers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go writer(i, &wg, sharedSlice)
	}

	// Launch readers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go reader(i, &wg, sharedSlice)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("Final slice contents:", sharedSlice.Get())
}
