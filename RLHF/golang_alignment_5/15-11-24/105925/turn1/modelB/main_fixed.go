package main

import (
	"fmt"
	"sync"
	"time"
)

// Database interface for write operations
type WriteDatabase interface {
	Write(key string, value int)
	GetData() map[string]int // Add a method to expose data
}

// WriteDatabase implementation using in-memory map
type writeDatabase struct {
	data map[string]int
	mu   sync.Mutex
}

func NewWriteDatabase() WriteDatabase {
	return &writeDatabase{data: make(map[string]int)}
}

func (db *writeDatabase) Write(key string, value int) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = value
}

// Getter method for accessing data
func (db *writeDatabase) GetData() map[string]int {
	db.mu.Lock()
	defer db.mu.Unlock()
	return db.data
}

// Database interface for read operations
type ReadDatabase interface {
	Read(key string) (int, bool)
}

// ReadDatabase implementation using in-memory map
type readDatabase struct {
	data map[string]int
}

func NewReadDatabase(data map[string]int) ReadDatabase {
	return &readDatabase{data: data}
}

func (db *readDatabase) Read(key string) (int, bool) {
	return db.data[key], true
}

func main() {
	// Simulate a large number of write operations
	writeDB := NewWriteDatabase()
	keys := make([]string, 10000)
	for i := range keys {
		keys[i] = fmt.Sprintf("key-%05d", i)
	}

	go func() {
		for i := 0; i < 10000; i++ {
			writeDB.Write(keys[i], i)
		}
	}()

	// Simulate delayed read operations
	time.Sleep(time.Second)

	// Use GetData to retrieve data from writeDB
	readDB := NewReadDatabase(writeDB.(*writeDatabase).GetData())

	// Perform a large number of read operations
	start := time.Now()
	for _, key := range keys {
		_, _ = readDB.Read(key)
	}
	elapsed := time.Since(start)
	fmt.Println("Read operation time:", elapsed)
}
