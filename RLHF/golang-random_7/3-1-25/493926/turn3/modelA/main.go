package main

import (
	"fmt"
	"sync"
	"time"
)

// A simple shared structure for demonstration purposes
type SharedData struct {
	Data map[string]int
	mx   sync.Mutex
}

func (s *SharedData) Read(key string) int {
	s.mx.Lock()
	defer s.mx.Unlock()

	value, exists := s.Data[key]
	if exists {
		return value
	}
	return 0
}

func (s *SharedData) Write(key string, value int) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.Data[key] = value
	fmt.Println("Wrote", value, "to key", key)
}

func main() {
	// Initialize the shared data
	data := &SharedData{Data: make(map[string]int)}

	// Create a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Generate key-value pairs for demonstration
	keys := []string{"a", "b", "c", "d"}
	for _, key := range keys {
		data.Write(key, 1)
	}

	// Start reading data concurrently
	numReaders := 10
	wg.Add(numReaders)
	for i := 0; i < numReaders; i++ {
		go func() {
			defer wg.Done()
			key := keys[i%len(keys)]
			value := data.Read(key)
			fmt.Printf("Read value %d for key %s\n", value, key)
			time.Sleep(time.Millisecond * 100)
		}()
	}

	// Start writing data concurrently
	numWriters := 5
	wg.Add(numWriters)
	for i := 0; i < numWriters; i++ {
		go func() {
			defer wg.Done()
			for {
				key := keys[i%len(keys)]
				value := data.Read(key) + 1
				data.Write(key, value)
				time.Sleep(time.Millisecond * 200)
			}
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All tasks completed.")
}
