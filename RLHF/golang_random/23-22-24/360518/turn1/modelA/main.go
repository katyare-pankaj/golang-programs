package main

import (
	"fmt"
	"time"
)

type Event struct {
	Type string
	Data interface{}
}

func main() {
	// Create a channel to pass events
	eventChannel := make(chan Event)

	// Start a goroutine to handle events of type "task1"
	go handleTask1(eventChannel)

	// Start a goroutine to handle events of type "task2"
	go handleTask2(eventChannel)

	// Simulate triggering events
	time.Sleep(2 * time.Second)
	eventChannel <- Event{"task1", "Data for task1"}

	time.Sleep(2 * time.Second)
	eventChannel <- Event{"task2", "Data for task2"}

	time.Sleep(2 * time.Second)
	eventChannel <- Event{"task1", "More data for task1"}

	// Close the event channel to signal the end
	close(eventChannel)
}

func handleTask1(eventChannel <-chan Event) {
	for event := range eventChannel {
		if event.Type == "task1" {
			fmt.Println("Handling task1 with data:", event.Data)
			time.Sleep(1 * time.Second) // Simulate some work
		}
	}
	fmt.Println("Task1 handler closed.")
}

func handleTask2(eventChannel <-chan Event) {
	for event := range eventChannel {
		if event.Type == "task2" {
			fmt.Println("Handling task2 with data:", event.Data)
			time.Sleep(2 * time.Second) // Simulate some work
		}
	}
	fmt.Println("Task2 handler closed.")
}
