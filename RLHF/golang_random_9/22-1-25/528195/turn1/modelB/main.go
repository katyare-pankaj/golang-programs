package main

import (
	"fmt"
	"time"
)

// LogEntry represents a single log entry
type LogEntry struct {
	Timestamp time.Time
	Message   string
}

// Logger generates log entries
type Logger struct {
	Entries chan LogEntry
}

// NewLogger creates a new logger
func NewLogger() *Logger {
	return &Logger{
		Entries: make(chan LogEntry, 100),
	}
}

// GenerateLogs generates log entries with a fixed interval
func (l *Logger) GenerateLogs() {
	for {
		l.Entries <- LogEntry{
			Timestamp: time.Now(),
			Message:   fmt.Sprintf("Log entry at %v", time.Now()),
		}
		time.Sleep(time.Second)
	}
}

func main() {
	// Create a new logger
	logger := NewLogger()
	// Start generating log entries in a goroutine
	go logger.GenerateLogs()

	// Use the range construct to iterate over log entries
	for entry := range logger.Entries {
		fmt.Printf("[%v] %s\n", entry.Timestamp, entry.Message)
		// Add your real-time analysis logic here
	}
}
