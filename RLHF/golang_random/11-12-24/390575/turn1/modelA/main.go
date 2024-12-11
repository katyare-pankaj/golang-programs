package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	mutex   sync.Mutex
	wg      sync.WaitGroup
)

func increment() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}

func main() {
	startTime := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()
	fmt.Printf("Mutex Counter: %d, Time: %v\n", counter, time.Since(startTime))
}
