package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numGoroutines = 1000
	iterations    = 100000
)

var (
	sharedCounter int
	mu            = sync.Mutex{}
	channel       = make(chan int)
)

func incrementWithMutex() {
	for i := 0; i < iterations; i++ {
		mu.Lock()
		sharedCounter++
		mu.Unlock()
	}
}

func incrementWithChannel() {
	for i := 0; i < iterations; i++ {
		channel <- 1
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Using mutexes
	startTime := time.Now()
	for i := 0; i < numGoroutines; i++ {
		go incrementWithMutex()
	}
	for i := 0; i < numGoroutines; i++ {
		go incrementWithMutex()
	}
	for i := 0; i < numGoroutines; i++ {
		go incrementWithMutex()
	}
	for i := 0; i < numGoroutines; i++ {
		go incrementWithMutex()
	}
	for i := 0; i < numGoroutines; i++ {
		go incrementWithMutex()
	}

	for i := 0; i < 5*numGoroutines; i++ {
		<-channel
	}
	fmt.Printf("Time with mutexes: %v\n", time.Since(startTime))
	fmt.Printf("Final counter value with mutexes: %d\n", sharedCounter)

	// Reset counter and channel
	sharedCounter = 0
	close(channel)
	channel = make(chan int)

	// Using channels
	startTime = time.Now()
	for i := 0; i < numGoroutines; i++ {
		go incrementWithChannel()
	}
	for i := 0; i < numGoroutines; i++ {
		go incrementWithChannel()
	}
	for i := 0; i < numGoroutines; i++ {
		go incrementWithChannel()
	}
	for i := 0; i < numGoroutines; i++ {
		go incrementWithChannel()
	}
	for i := 0; i < numGoroutines; i++ {
		go incrementWithChannel()
	}

	for i := 0; i < 5*numGoroutines; i++ {
		<-channel
	}
	fmt.Printf("Time with channels: %v\n", time.Since(startTime))
	fmt.Printf("Final counter value with channels: %d\n", sharedCounter)
}
