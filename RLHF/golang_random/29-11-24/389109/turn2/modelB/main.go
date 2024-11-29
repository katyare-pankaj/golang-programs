package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// SharedData represents our shared resource
type SharedData struct {
	data    map[int]int
	rwMutex sync.RWMutex
}

func newSharedData() *SharedData {
	return &SharedData{
		data: make(map[int]int),
	}
}

// Read access, multiple readers allowed
func (sd *SharedData) Read(key int) int {
	sd.rwMutex.RLock()
	defer sd.rwMutex.RUnlock()
	return sd.data[key]
}

// Write access, only one writer allowed
func (sd *SharedData) Write(key, value int) {
	sd.rwMutex.Lock()
	defer sd.rwMutex.Unlock()
	sd.data[key] = value
}

func main() {
	var wg sync.WaitGroup
	const numThreads = 10
	const numIterations = 10000
	sharedData := newSharedData()

	wg.Add(numThreads)

	// launching several readers
	for i := 0; i < numThreads; i++ {
		go func() {
			for j := 0; j < numIterations; j++ {
				key := rand.Intn(100)
				value := sharedData.Read(key)
				if value == 0 {
					sharedData.Write(key, j) //writing new value
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// Display the contents of the shared data
	fmt.Println("Final Shared Data:")
	sharedData.rwMutex.RLock()
	for k, v := range sharedData.data {
		fmt.Printf("%d: %d\n", k, v)
	}
	sharedData.rwMutex.RUnlock()
}
