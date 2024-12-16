package main

import (
	"fmt"
	"sync"
)

var count int
var mu sync.Mutex
var wg sync.WaitGroup

func increment() {
	for i := 0; i < 100000; i++ {
		mu.Lock()
		count++
		mu.Unlock()
	}
	wg.Done()
}

func decrement() {
	for i := 0; i < 100000; i++ {
		mu.Lock()
		count--
		mu.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	go increment()
	go decrement()

	wg.Wait()

	fmt.Println("Final count:", count) // Should always output 0
}
