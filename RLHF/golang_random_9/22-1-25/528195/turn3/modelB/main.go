package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

type LogEntry struct {
	Timestamp time.Time
	Message   string
}

const (
	numWorkers    = 4          // Number of worker goroutines
	logBufferSize = 1000       // Buffer size for log entries
	quitSignal    = struct{}{} // Empty struct to signal goroutines to quit
)

func logWorker(workerID int, logs <-chan LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", workerID)
	for logEntry := range logs {
		// Simulate processing the log entry
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Printf("[Worker %d] [%s] %s\n", workerID, logEntry.Timestamp.Format(time.RFC3339), logEntry.Message)
	}
	fmt.Printf("Worker %d finished\n", workerID)
}

func simulateLogEntries(logs chan<- LogEntry, quit <-chan struct{}) {
	defer close(logs)
	messages := []string{
		"Started processing request",
		"Connecting to database",
		"Fetching data",
		"Sending response",
		"Request completed",
	}
	for {
		select {
		case <-quit:
			return // Return when quit signal is received
		default:
			for _, msg := range messages {
				logs <- LogEntry{Timestamp: time.Now(), Message: msg}
			}
			time.Sleep(time.Second) // Simulate time between log entries
		}
	}
}

func main() {
	logs := make(chan LogEntry, logBufferSize)
	quit := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go logWorker(i, logs, &wg)
	}

	go simulateLogEntries(logs, quit)

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All log entries processed.")
	close(quit) // Signal the simulateLogEntries goroutine to quit
}
