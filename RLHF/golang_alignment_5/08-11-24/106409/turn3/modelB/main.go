package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SharedData represents a structure that holds shared data.
type SharedData struct {
	value int
	mutex sync.RWMutex
}

// Read reads the shared data value safely.
func (sd *SharedData) Read() int {
	sd.mutex.RLock()
	defer sd.mutex.RUnlock()
	return sd.value
}

// Write updates the shared data value safely.
func (sd *SharedData) Write(value int) {
	sd.mutex.Lock()
	defer sd.mutex.Unlock()
	sd.value = value
}

// Inc increments the shared data value safely.
func (sd *SharedData) Inc() {
	sd.mutex.Lock()
	defer sd.mutex.Unlock()
	sd.value++
}

// Dec decrements the shared data value safely.
func (sd *SharedData) Dec() {
	sd.mutex.Lock()
	defer sd.mutex.Unlock()
	sd.value--
}

// accessSharedData simulates concurrent read and write access to the shared data.
func accessSharedData(wg *sync.WaitGroup, sd *SharedData) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		// Simulate read access
		currentValue := sd.Read()
		fmt.Printf("Read: %d\n", currentValue)

		// Simulate write access (increment or decrement randomly)
		if rand.Intn(2) == 0 {
			sd.Inc()
		} else {
			sd.Dec()
		}
	}
}

func main() {
	const numConcurrentAccesses = 5

	sharedData := &SharedData{value: 0}

	var wg sync.WaitGroup
	wg.Add(numConcurrentAccesses)

	for i := 0; i < numConcurrentAccesses; i++ {
		go accessSharedData(&wg, sharedData)
	}

	wg.Wait()
	fmt.Println("Final shared data value:", sharedData.Read())
}
