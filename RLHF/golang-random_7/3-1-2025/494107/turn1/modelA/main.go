package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func sayHello(w *sync.WaitGroup, message string) {
	defer w.Done()          // Signal to WaitGroup that this task is complete
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Hello, %s!\n", message)
}

func main() {
	// Create a new WaitGroup
	var wg sync.WaitGroup

	// Start three concurrent hello tasks
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the counter
		go sayHello(&wg, fmt.Sprintf("world%d", i))
	}

	// Wait for all concurrent tasks to complete
	wg.Wait()

	fmt.Println("All tasks completed.")
}

func TestSayHello(t *testing.T) {
	// Simulating tasks with timers
	expectedOutput := []string{"Hello, world1!", "Hello, world2!", "Hello, world3!"}
	receivedOutput := []string{}

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(message string) {
			defer wg.Done()
			time.Sleep(time.Second)
			receivedOutput = append(receivedOutput, message)
		}(fmt.Sprintf("world%d", i))
	}

	wg.Wait()

	if !assertEqual(expectedOutput, receivedOutput) {
		t.Errorf("Expected: %v, Received: %v", expectedOutput, receivedOutput)
	}
}

func assertEqual(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
