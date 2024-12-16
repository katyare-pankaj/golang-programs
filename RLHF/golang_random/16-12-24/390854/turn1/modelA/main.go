package main

import (
	"fmt"
	"sync"
)

var counter int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter++
	}
}

func main() {
	var wg sync.WaitGroup

	// Start 10 goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Final counter value:", counter)
}
