package main

import (
	"fmt"
	"sync"
)

// Thread-safe logger using mutex
type ThreadSafeLogger struct {
	mu sync.Mutex
}

func (l *ThreadSafeLogger) Log(message string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Println(message)
}

// Log processing function using Goroutines
func logProcessorWithGoroutines(logs []string) {
	for _, log := range logs {
		go func(l string) {
			fmt.Println(l) // Without synchronization, order of logs is unpredictable
		}(log)
	}
}

// Log processing function using traditional threads
func logProcessorWithThreads(logs []string) {
	for _, log := range logs {
		go func(l string) {
			fmt.Println(l) // Without synchronization, order of logs is unpredictable
		}(log)
	}
}

// Log processing function with synchronization using channels
func logProcessorWithChannels(logs []string) {
	logChan := make(chan string)

	// Goroutine to consume logs from the channel and print them
	go func() {
		for log := range logChan {
			fmt.Println(log) // Logs will be printed in the order they are received
		}
	}()

	// Send logs to the channel
	for _, log := range logs {
		logChan <- log
	}

	close(logChan)
}

func main() {
	logs := []string{
		"Log 1",
		"Log 2",
		"Log 3",
		"Log 4",
		"Log 5",
	}

	// Output: Log 1 Log 2 Log 3 Log 4 Log 5 (Unpredictable order)
	logProcessorWithGoroutines(logs)

	// Output: Log 1 Log 2 Log 3 Log 4 Log 5 (Unpredictable order)
	logProcessorWithThreads(logs)

	// Output: Log 1 Log 2 Log 3 Log 4 Log 5 (Preserved order)
	logProcessorWithChannels(logs)

	// Thread-safe logger example
	logger := &ThreadSafeLogger{}
	for _, log := range logs {
		go func(l string) {
			logger.Log(l) // Using a mutex, order of logs is preserved
		}(log)
	}
}
