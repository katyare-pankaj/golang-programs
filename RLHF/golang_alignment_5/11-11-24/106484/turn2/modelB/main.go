package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// Code represents the shared code for collaborative editing
type Code struct {
	content string
	version uint64 // Use uint64 for version, as atomic operations support only unsigned integers
}

// updateCode uses atomic operations to update the code and version number
func updateCode(code *Code, updates chan string) {
	for update := range updates {
		for {
			// Read the current version before updating
			currentVersion := atomic.LoadUint64(&code.version)

			// Perform the update only if no other thread has modified the code since reading the version
			if atomic.CompareAndSwapUint64(&code.version, currentVersion, currentVersion+1) {
				code.content += update + "\n"
				fmt.Println("Updated code:", code.content)
				break
			}
			// If another thread updated the code, try again (spinlock approach, not ideal in high contention scenarios)
			// In a real application, you could use a more advanced mechanism like channel communication or a mutex with wait groups.
			time.Sleep(1 * time.Millisecond)
		}
	}
}

func main() {
	// Initialize the shared code and version
	sharedCode := &Code{content: "", version: 0}
	updates := make(chan string)

	// Start the goroutine to handle code updates using atomic operations
	go updateCode(sharedCode, updates)

	// Simulate multiple developers making changes
	developer1 := "Alice"
	developer2 := "Bob"

	go func() {
		updates <- fmt.Sprintf("%s added a new function", developer1)
		updates <- fmt.Sprintf("%s changed variable 'x' to 'y'", developer1)
	}()

	go func() {
		updates <- fmt.Sprintf("%s fixed a typo in comment", developer2)
		updates <- fmt.Sprintf("%s added a new line at the end", developer2)
	}()

	// Wait for all updates to complete
	time.Sleep(2 * time.Second)

	// Output the final state of the code and version
	fmt.Println("Final Code:")
	fmt.Println(sharedCode.content)
	fmt.Println("Final Version:", atomic.LoadUint64(&sharedCode.version))
}
