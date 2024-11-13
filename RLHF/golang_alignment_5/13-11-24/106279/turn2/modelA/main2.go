package main

import (
	"fmt"
	"math/rand"

	"sync"
	"time"
)

// counter is a shared resource that will be incremented by multiple goroutines
var counter int
var mutex sync.Mutex

// incrementCounter increments the global counter by 1 after a random delay
func incrementCounter() {
	mutex.Lock()
	defer mutex.Unlock()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	counter++
}

func main() {
	const numGoroutines = 10
	const numIncrements = 100

	// Start multiple goroutines that will increment the counter
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < numIncrements; j++ {
				incrementCounter()
			}
		}()
	}

	// Wait for all goroutines to finish
	time.Sleep(time.Second * 2)

	// Print the final value of the counter, which should be the sum of all increments
	fmt.Println("Final counter value:", counter)
}
