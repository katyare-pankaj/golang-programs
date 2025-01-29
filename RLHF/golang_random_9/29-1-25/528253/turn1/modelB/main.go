package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func eventLoop(eventCh chan string) {
	defer close(eventCh) // Close the event channel when the loop exits
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a message: ")
		text, _ := reader.ReadString('\n')
		eventCh <- text // Send the event to the event channel
	}
}

func main() {
	eventCh := make(chan string)
	defer close(eventCh)

	go eventLoop(eventCh) // Start the event loop in a goroutine

	for {
		select {
		case event := <-eventCh: // Receive events from the channel
			fmt.Println("Received event:", event)
		default: // Non-blocking select to check for channel closure
			if eventCh == nil {
				fmt.Println("Event loop terminated.")
				return
			}
			// Add a small delay to avoid busy waiting
			time.Sleep(100 * time.Millisecond)
		}
	}
}
