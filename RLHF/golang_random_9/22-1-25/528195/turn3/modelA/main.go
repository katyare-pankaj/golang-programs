package main

import (
	"fmt"
	"sync"
	"time"
)

// LogEntry represents a single log entry.
type LogEntry struct {
	Timestamp time.Time
	Message   string
}

const (
	numWorkers = 4 // Number of goroutines (workers) to process log entries
)

// simulateLogEntries simulates sending log entries to the logs channel.
func simulateLogEntries(logs chan<- LogEntry) {
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

// processLogEntries processes log entries from the logs channel.
func processLogEntries(workerId int, logs <-chan LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	for logEntry := range logs {
		// Process the log entry (e.g., print it)
		fmt.Printf("Worker %d: [%s] %s\n", workerId, logEntry.Timestamp.Format(time.RFC3339), logEntry.Message)
	}
}

func main() {
	// Create a channel to receive log entries
	logs := make(chan LogEntry)

	// WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go processLogEntries(i, logs, &wg)
	}

	// Start a goroutine to simulate log entries
	go simulateLogEntries(logs)

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All log entries processed.")
}
