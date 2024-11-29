package main

import (
	"fmt"
	"sync"
	"time"
)

type SharedResource struct {
	data []int
	rwmu sync.RWMutex
}

func (s *SharedResource) Read() {
	s.rwmu.RLock()         // Acquire a read lock
	defer s.rwmu.RUnlock() // Release the read lock
	fmt.Println("Current data:", s.data)
}

func (s *SharedResource) Write(newData []int) {
	s.rwmu.Lock()         // Acquire a write lock
	defer s.rwmu.Unlock() // Release the write lock
	s.data = newData
	fmt.Println("Data updated to:", s.data)
}

func main() {
	var sr SharedResource
	sr.data = []int{1, 2, 3, 4, 5}
	const numReadThreads = 10
	const numWriteThreads = 2
	const numIterations = 10

	// Function for a reader thread
	readThread := func() {
		for i := 0; i < numIterations; i++ {
			sr.Read()
			time.Sleep(100 * time.Microsecond)
		}
	}

	// Function for a writer thread
	writeThread := func() {
		for i := 0; i < numIterations; i++ {
			newData := sr.data[:]
			newData[0] += i + 1
			sr.Write(newData)
			time.Sleep(500 * time.Microsecond)
		}
	}

	// Start reader threads
	for i := 0; i < numReadThreads; i++ {
		go readThread()
	}

	// Start writer threads
	for i := 0; i < numWriteThreads; i++ {
		go writeThread()
	}

	// Wait for all threads to complete
	time.Sleep(2 * time.Second)
}
