package main

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

// you have a function that needs to be tested concurrently
func worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()                // Decrement the wait group count on return
	num := <-ch                    // Receive a value from the channel
	result := num * num            // Perform some operation
	fmt.Println("Result:", result) // Output the result
}
func TestWorkerConcurrency(t *testing.T) {
	const numWorkers = 3
	var wg sync.WaitGroup
	ch := make(chan int)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)          // Increment the wait group count
		go worker(ch, &wg) // Launch a goroutine for each worker
	}
	// Send some data to the channel
	for i := 0; i < numWorkers; i++ {
		ch <- i + 1
	}
	close(ch) // Close the channel to signal the end of data
	wg.Wait() // Wait for all goroutines to finish
	// Verify the output (with a simple assert)
	expectedOutput := []string{
		"Result: 1",
		"Result: 4",
		"Result: 9",
	}
	// Reset testing output buffer
	t.Log("Expected Output:")
	for _, output := range expectedOutput {
		t.Log(output)
	}
	// Comparison of outputs
	gotOutput := testing.CaptureStdout(func() {
		fmt.Println("Here is the normal output:")
		t.Logf("Test completed.")
	})
	t.Log("Got Output:")
	t.Log(gotOutput)
	// Test for equal strings
	for _, output := range expectedOutput {
		if !isContained(gotOutput, output) {
			t.Fatalf("Missing expected output: %s", output)
		}
	}
}

func isContained(str, subStr string) bool {
	return strings.Contains(str, subStr)
}
