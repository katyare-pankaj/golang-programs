package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var counter int
	var mutex sync.Mutex

	// Start 100 goroutines that will increment the counter
	for i := 0; i < 100; i++ {
		go incrementCounter(&counter, &mutex)
	}

	// Wait for all goroutines to finish
	// (In a real application, use a sync.WaitGroup instead)
	time.Sleep(1 * time.Second)

	// Print the final value of the counter
	fmt.Println("Final value of the counter:", counter)
}

func incrementCounter(counter *int, mutex *sync.Mutex) {
	for i := 0; i < 10000; i++ {
		mutex.Lock()
		*counter++
		mutex.Unlock()
	}
}
