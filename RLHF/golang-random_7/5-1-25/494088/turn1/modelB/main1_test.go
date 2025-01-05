package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

const (
	numWorkers = 100
	numJobs    = 10000
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numJobs/numWorkers; i++ {
		time.Sleep(1 * time.Millisecond) // Simulate work
	}
}

func TestWaitGroupReliability(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(&wg)
	}
	wg.Wait()

	// Ensure no unexpected allocations (sanity check)
	if testing.AllocsPerRun(1000, func() {
		wg.Wait()
	}) != 0 {
		t.Error("WaitGroup did not correctly wait for all goroutines to finish")
	}
}

func TestWaitGroupTimeout(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		if i == numWorkers-1 {
			// One worker intentionally delays indefinitely
			go func() {
				defer wg.Done()
				select {} // Block indefinitely
			}()
		} else {
			go worker(&wg)
		}
	}

	// Use a timeout mechanism with a channel
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		t.Error("WaitGroup completed unexpectedly; expected timeout")
	case <-time.After(2 * time.Second):
		// Test succeeded with timeout
		fmt.Println("Timeout as expected")
	}
}
