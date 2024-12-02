package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	// Shared data
	conditionMet := false

	// Goroutine waiting for the condition
	go func() {
		mu.Lock()
		for !conditionMet {
			cond.Wait()
		}
		fmt.Println("Condition variable met, proceeding.")
		mu.Unlock()
	}()

	// Simulate some work before setting the condition
	time.Sleep(2 * time.Second)

	// Goroutine that sets the condition
	go func() {
		mu.Lock()
		conditionMet = true
		cond.Signal() // Could also use cond.Broadcast() to wake all waiting goroutines
		fmt.Println("Condition variable signaled.")
		mu.Unlock()
	}()

	time.Sleep(1 * time.Second)
}
