package main

import (
	"fmt"
	"sync"
	"time"
)

type SharedData struct {
	ready bool
	cond  *sync.Cond
}

func main() {
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)

	shared := SharedData{
		ready: false,
		cond:  cond,
	}

	// Goroutine that waits for the condition
	go func() {
		shared.cond.L.Lock()
		for !shared.ready {
			shared.cond.Wait()
		}
		fmt.Println("Condition met, proceeding.")
		shared.cond.L.Unlock()
	}()

	// Simulate some work before the condition is met
	time.Sleep(2 * time.Second)

	// Goroutine that sets the condition
	go func() {
		shared.cond.L.Lock()
		shared.ready = true
		shared.cond.Signal() // can use Broadcast() to wake all waiting goroutines
		fmt.Println("Condition set, signaled.")
		shared.cond.L.Unlock()
	}()

	// Give some time for goroutines to finish
	time.Sleep(1 * time.Second)
}
