package main

import (
	"fmt"
	"sync"
	"time"
)

// Define the Table struct
type Table struct {
	data  map[string]int
	mutex sync.Mutex
}

// NewTable creates a new table
func NewTable() *Table {
	return &Table{data: make(map[string]int)}
}

// Set updates the value in the table
func (t *Table) Set(key string, value int) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.data[key] = value
}

// Get retrieves the value from the table
func (t *Table) Get(key string) (int, bool) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.data[key], t.data[key] != 0
}

// Example usage
func main() {
	table := NewTable()

	var wg sync.WaitGroup

	// Create goroutines for concurrent operations
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				key := fmt.Sprintf("key%d_%d", i, j)
				value := j + i*10
				table.Set(key, value)
				time.Sleep(time.Millisecond * 100)

				// Read data to simulate real-time collaboration
				readKey := fmt.Sprintf("key%d_%d", (i+1)%5, j)
				readValue, exists := table.Get(readKey)
				if exists {
					fmt.Printf("Goroutine %d: Read key %s -> %d\n", i, readKey, readValue)
				} else {
					fmt.Printf("Goroutine %d: Key %s not found\n", i, readKey)
				}
			}
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Print the final state of the table
	fmt.Println("Final state of the table:")
	for key, value := range table.data {
		fmt.Printf("%s: %d\n", key, value)
	}
}
