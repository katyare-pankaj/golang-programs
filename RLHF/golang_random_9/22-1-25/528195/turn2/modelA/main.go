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
		logs <- LogEntry{
			Timestamp: time.Now(),
			Message:   msg,
		}
		time.Sleep(500 * time.Millisecond) // Simulate time between log entries
	}
	close(logs) // Close the channel when done sending log entries
}

// processLog is a worker function that processes log entries.
func processLog(id int, logs <-chan LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	for logEntry := range logs {
		// Simulating log processing
		fmt.Printf("Worker %d processing log: [%s] %s\n", id, logEntry.Timestamp.Format(time.RFC3339), logEntry.Message)
		time.Sleep(1 * time.Second) // Simulate processing time
	}
}

func main() {
	logs := make(chan LogEntry, 10) // Buffered channel to hold log entries
	const numWorkers = 3            // Number of workers

	// WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go processLog(i, logs, &wg)
	}

	// Simulate log entries being added
	go simulateLogEntries(logs)

	// Wait for all workers to complete processing
	wg.Wait()

	fmt.Println("All log entries processed.")
}
