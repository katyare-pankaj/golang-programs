package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var sharedData struct {
		value int
		cond  *sync.Cond
	}
	sharedData.cond = sync.NewCond(&sync.Mutex{})

	// Goroutine that waits for the value to be 10
	go func() {
		sharedData.cond.L.Lock()
		for sharedData.value != 10 {
			sharedData.cond.Wait()
		}
		fmt.Println("Value is now 10, proceeding.")
		sharedData.cond.L.Unlock()
	}()

	// Goroutine that sets the value to 10
	go func() {
		sharedData.cond.L.Lock()
		time.Sleep(2 * time.Second) // Simulate some work
		sharedData.value = 10
		sharedData.cond.Signal()
		fmt.Println("Value set to 10, signaling.")
		sharedData.cond.L.Unlock()
	}()

	// Wait for both goroutines to finish
	time.Sleep(3 * time.Second)
}
