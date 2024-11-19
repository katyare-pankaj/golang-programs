package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int = 0
	mu          = &sync.Mutex{}
)

func increment(num int) {
	for i := 0; i < num; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func main() {
	const numGoroutines = 10
	const iterations = 100000

	for i := 0; i < numGoroutines; i++ {
		go increment(iterations)
	}

	time.Sleep(2 * time.Second)

	fmt.Printf("Final counter value: %d\n", counter)
}
