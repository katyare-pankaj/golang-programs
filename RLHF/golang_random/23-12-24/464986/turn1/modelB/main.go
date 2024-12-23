package main

import (
	"fmt"
	"time"
)

// Event struct to store event data
type Event struct {
	Name        string
	Time        time.Time
	Description string
}

func main() {
	// Create a slice to store events
	var events []Event

	// Loop to add events
	for {
		var name, desc string
		var eventTime time.Time

		fmt.Print("Enter event name (or 'exit' to quit): ")
		fmt.Scanln(&name)

		if name == "exit" {
			break
		}
		fmt.Print("Enter event time (YYYY-MM-DD HH:MM): ")
		_, err := fmt.Scanln(&eventTime)
		if err != nil {
			fmt.Println("Invalid time format. Please try again.")
			continue
		}

		fmt.Print("Enter event description: ")
		fmt.Scanln(&desc)

		// Create a new event and append it to the slice
		newEvent := Event{Name: name, Time: eventTime, Description: desc}
		events = append(events, newEvent)
	}

	// Display all events
	fmt.Println("\nScheduled Events:")
	for _, event := range events {
		fmt.Printf("Event Name: %s\nEvent Time: %s\nEvent Description: %s\n",
			event.Name, event.Time.Format("2006-01-02 15:04"), event.Description)
		fmt.Println("--------------------")
	}
}
