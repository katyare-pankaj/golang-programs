package main

import (
	"fmt"
	"time"
)

// LogEntry represents a single log entry.
type LogEntry struct {
	Timestamp time.Time
	Message   string
}

// simulateLogEntries simulates sending log entries to the logs channel.
func simulateLogEntries(logs chan<- LogEntry) {
	// Simulating some log entries
	messages := []string{
		"Started processing request",
		"Connecting to database",
		"Fetching data",
		"Sending response",
		"Request completed",
	}

	for _, msg := range messages {
		// Send a log entry to the channel
		logs <- LogEntry{
			Timestamp: time.Now(),
			Message:   msg,
		}
		time.Sleep(1 * time.Second) // Simulate time between log entries
	}

	// Close the channel after sending all log entries
	close(logs)
}

func main() {
	// Create a channel to receive log entries
	logs := make(chan LogEntry)

	// Start a goroutine to simulate log entries
	go simulateLogEntries(logs)

	// Use range to iterate over log entries from the channel
	for logEntry := range logs {
		// Process the log entry (in this case, just print it)
		fmt.Printf("[%s] %s\n", logEntry.Timestamp.Format(time.RFC3339), logEntry.Message)
	}

	fmt.Println("All log entries processed.")
}
