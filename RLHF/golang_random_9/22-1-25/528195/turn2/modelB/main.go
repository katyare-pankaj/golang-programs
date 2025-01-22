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
	logEntryBufferSize     = 1000 // Adjust this buffer size based on your needs
	numProcessorGoroutines = 4    // Number of goroutines to process log entries
)

func simulateLogEntries(logs chan<- LogEntry) {
	messages := []string{
		"Started processing request",
		"Connecting to database",
		"Fetching data",
		"Sending response",
		"Request completed",
	}

	for _, msg := range messages {
		logs <- LogEntry{Timestamp: time.Now(), Message: msg}
		time.Sleep(1 * time.Millisecond) // Simulate time between log entries
	}
	close(logs)
}

func processLogEntry(logEntry LogEntry) {
	// Simulate processing the log entry
	time.Sleep(50 * time.Millisecond)
	// Replace this with your actual log processing logic
	fmt.Printf("[%s] %s\n", logEntry.Timestamp.Format(time.RFC3339), logEntry.Message)
}

func processLogEntries(logs <-chan LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	for logEntry := range logs {
		processLogEntry(logEntry)
	}
}

func main() {
	logs := make(chan LogEntry, logEntryBufferSize)
	var wg sync.WaitGroup

	wg.Add(numProcessorGoroutines)
	for i := 0; i < numProcessorGoroutines; i++ {
		go processLogEntries(logs, &wg)
	}

	go simulateLogEntries(logs)

	wg.Wait()
	fmt.Println("All log entries processed.")
}
