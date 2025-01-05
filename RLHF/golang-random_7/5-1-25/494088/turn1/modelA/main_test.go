package main

import (
	"sync"
	"testing"
	"time"
)

func TestWaitGroupReliability(t *testing.T) {
	const (
		numGoroutines = 1000
		iterations    = 10000
		duration      = time.Millisecond * 500 // 500 ms
	)

	wg := sync.WaitGroup{}

	// Function to simulate a task
	task := func(id int) {
		defer wg.Done()
		time.Sleep(duration)
	}

	// Start the goroutines
	for i := 0; i < numGoroutines; i++ {
		for j := 0; j < iterations; j++ {
			wg.Add(1)
			go task(i*iterations + j)
		}
	}

	// Wait for all goroutines to complete
	start := time.Now()
	wg.Wait()
	end := time.Now()

	// Calculate total execution time
	totalTime := end.Sub(start)

	// Output the result
	t.Logf("Total time for %d goroutines and %d iterations: %v", numGoroutines, iterations, totalTime)
}

func TestWaitGroupRaceCondition(t *testing.T) {
	const (
		numGoroutines = 1000
		iterations    = 10000
		duration      = time.Millisecond * 500 // 500 ms
	)

	wg := sync.WaitGroup{}

	// Function to simulate a task
	task := func(id int) {
		defer wg.Done()
		time.Sleep(duration)
	}

	// Start the goroutines
	for i := 0; i < numGoroutines; i++ {
		for j := 0; j < iterations; j++ {
			// Deliberately skip Add() to simulate a race condition
			go task(i*iterations + j)
		}
	}

	// Wait for all goroutines to complete
	start := time.Now()
	wg.Wait()
	end := time.Now()

	// Calculate total execution time
	totalTime := end.Sub(start)

	// Output the result
	t.Logf("Total time for %d goroutines and %d iterations with race condition: %v", numGoroutines, iterations, totalTime)

	// Assert that all goroutines completed correctly
	// This can be tricky to assert directly in Go due to the nature of race conditions
	// but can be checked by observing the output or using external tools like race detector
}

func TestWaitGroupMemoryLeak(t *testing.T) {
	const (
		numGoroutines = 1000
		iterations    = 10000
		duration      = time.Millisecond * 500 // 500 ms
	)

	// Create a WaitGroup that is never used
	wg := sync.WaitGroup{}

	// Function to simulate a task
	task := func(id int) {
		defer wg.Done()
		time.Sleep(duration)
	}

	// Start the goroutines
	for i := 0; i < numGoroutines; i++ {
		for j := 0; j < iterations; j++ {
			// Deferring Done() on a WaitGroup that is never Add()'ed will lead to memory leaks
			go task(i*iterations + j)
		}
	}

	// Wait for all goroutines to complete
	start := time.Now()
	wg.Wait()
	end := time.Now()

	// Calculate total execution time
	totalTime := end.Sub(start)

	// Output the result
	t.Logf("Total time for %d goroutines and %d iterations with memory leak: %v", numGoroutines, iterations, totalTime)

	// Check for memory leaks using tools like `go test -race` or Valgrind
}

func BenchmarkWaitGroupPerformance(b *testing.B) {
	const (
		numGoroutines = 1000
		iterations    = 10000
		duration      = time.Millisecond * 500 // 500 ms
	)

	wg := sync.WaitGroup{}

	// Function to simulate a task
	task := func(id int) {
		defer wg.Done()
		time.Sleep(duration)
	}

	// Start the goroutines
	for i := 0; i < numGoroutines; i++ {
		for j := 0; j < iterations; j++ {
			wg.Add(1)
			go task(i*iterations + j)
		}
	}

	// Wait for all goroutines to complete
	wg.Wait()
}
