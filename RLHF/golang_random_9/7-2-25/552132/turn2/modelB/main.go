package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	numGoroutines  = 10
	numOperations  = 100000
	sharedDataSize = 10
)

// SharedData represents a shared data structure that can be modified by multiple goroutines.
type SharedData struct {
	mutex sync.RWMutex // Mutex for synchronizing access
	data  []int
}

func (sd *SharedData) init() {
	sd.data = make([]int, sharedDataSize)
}

func (sd *SharedData) read() int {
	sd.mutex.RLock()
	defer sd.mutex.RUnlock()
	// Choose a random index and return the value
	idx := rand.Intn(len(sd.data))
	return sd.data[idx]
}

func (sd *SharedData) write(value int) {
	sd.mutex.Lock()
	defer sd.mutex.Unlock()
	// Choose a random index and update the value
	idx := rand.Intn(len(sd.data))
	sd.data[idx] = value
}

func worker(sharedData *SharedData, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numOperations; i++ {
		// Perform random read or write operations
		if rand.Intn(2) == 0 {
			sharedData.write(rand.Intn(1000))
		} else {
			_ = sharedData.read()
		}
	}
}

func main() {
	var wg sync.WaitGroup
	sharedData := SharedData{}
	sharedData.init()

	// Start multiple goroutines
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go worker(&sharedData, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Display the final state of the shared data
	fmt.Println("Final state of shared data:", sharedData.data)
}
