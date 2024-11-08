package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SensitiveData represents a class holding sensitive data
type SensitiveData struct {
	Data  []byte
	mutex sync.RWMutex
}

// NewSensitiveData creates a new SensitiveData instance
func NewSensitiveData(data []byte) *SensitiveData {
	return &SensitiveData{Data: data}
}

// Read reads the sensitive data safely
func (sd *SensitiveData) Read() []byte {
	sd.mutex.RLock()
	defer sd.mutex.RUnlock()
	return sd.Data
}

// Update updates the sensitive data safely
func (sd *SensitiveData) Update(data []byte) {
	sd.mutex.Lock()
	defer sd.mutex.Unlock()
	sd.Data = data
}

// accessSensitiveData simulates concurrent access to sensitive data
func accessSensitiveData(wg *sync.WaitGroup, data *SensitiveData) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		// Simulate read access
		readData := data.Read()
		fmt.Printf("Read: %s\n", readData)

		// Simulate update access
		newData := []byte(fmt.Sprintf("Updated data %d", i))
		data.Update(newData)
	}
}

func main() {
	const sensitiveDataSize = 32
	sensitiveData := NewSensitiveData(make([]byte, sensitiveDataSize))
	rand.Seed(time.Now().UnixNano())

	// Generate random initial data
	for i := range sensitiveData.Data {
		sensitiveData.Data[i] = byte(rand.Intn(256))
	}

	var wg sync.WaitGroup
	numConcurrentAccesses := 5
	wg.Add(numConcurrentAccesses)

	for i := 0; i < numConcurrentAccesses; i++ {
		go accessSensitiveData(&wg, sensitiveData)
	}

	wg.Wait()
	fmt.Println("Sensitive data processing completed")
}
