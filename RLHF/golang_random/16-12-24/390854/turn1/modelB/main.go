package main

import (
	"fmt"
	"sync"
)

var count int
var wg sync.WaitGroup

func increment() {
	for i := 0; i < 100000; i++ {
		count++
	}
	wg.Done()
}

func decrement() {
	for i := 0; i < 100000; i++ {
		count--
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	go increment()
	go decrement()

	wg.Wait()

	fmt.Println("Final count:", count)
	// Expected output should be 0, but without synchronization, it might not be.
}
