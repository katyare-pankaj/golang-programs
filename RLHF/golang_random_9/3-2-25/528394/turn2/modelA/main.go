package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	mtx     sync.Mutex
)

func incrementWithoutMutex(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		value := counter
		// Simulate some processing
		time.Sleep(time.Millisecond)
		counter = value + 1
		fmt.Printf("Without Mutex: Goroutine %d incremented counter to %d\n", id, counter)
	}
}

func incrementWithMutex(id int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		mtx.Lock() // Lock the mutex before accessing the shared counter
		counter++
		fmt.Printf("With Mutex: Goroutine %d incremented counter to %d\n", id, counter)
		mtx.Unlock()                 // Unlock the mutex after accessing the shared counter
		time.Sleep(time.Millisecond) // Simulate some processing
	}
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 5

	// Part 1: Without Mutex
	fmt.Println("Incrementing counter without Mutex:")
	counter = 0
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrementWithoutMutex(i, &wg)
	}
	wg.Wait()
	fmt.Printf("Final counter value without Mutex: %d\n\n", counter)

	// Part 2: With Mutex
	fmt.Println("Incrementing counter with Mutex:")
	counter = 0
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrementWithMutex(i, &wg, &mtx)
	}
	wg.Wait()
	fmt.Printf("Final counter value with Mutex: %d\n", counter)
}
