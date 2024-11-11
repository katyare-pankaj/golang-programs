package main

import (
	"fmt"
	"sync"
)

// Define a Code struct to represent the shared code
type Code struct {
	sync.Mutex
	content string
}

// Function to update the code content
func updateCode(code *Code, updates chan string) {
	for update := range updates {
		code.Lock()
		code.content += update + "\n"
		code.Unlock()
		fmt.Println("Updated code:", code.content)
	}
}

func main() {
	// Create a shared Code instance
	sharedCode := &Code{}

	// Create a channel for real-time code updates
	updates := make(chan string)

	// Start a goroutine to handle code updates
	go updateCode(sharedCode, updates)

	// Simulate multiple developers collaborating
	developer1 := "Alice"
	developer2 := "Bob"

	go func() {
		updates <- fmt.Sprintf("%s added a new line", developer1)
		updates <- fmt.Sprintf("%s commented: 'Fixing bug in loop'", developer1)
	}()

	go func() {
		updates <- fmt.Sprintf("%s made a change to function 'calc'", developer2)
		updates <- fmt.Sprintf("%s suggested using 'ioutil' for file handling", developer2)
	}()

	// Main thread waits for updates to complete
	// In a real app, you could add a way for developers to signal completion or use a timeout
	var input string
	fmt.Scanln(&input)
}
