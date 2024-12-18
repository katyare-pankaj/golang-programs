package main

import (
	"fmt"
	"sync"
	"testing"
)

// worker is a simple function that processes a number and prints it.
func worker(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	fmt.Printf("Processing %d\n", num)
}

// TestWorker demonstrates how to use a sync.WaitGroup to test Goroutines.
func TestWorker(t *testing.T) {
	var wg sync.WaitGroup
	numbers := []int{1, 2, 3, 4, 5}

	for _, num := range numbers {
		wg.Add(1)
		go worker(&wg, num)
	}

	wg.Wait()
}

func TestMain(m *testing.M) {
	// Call the testing framework's "main"
	m.Run()
}
