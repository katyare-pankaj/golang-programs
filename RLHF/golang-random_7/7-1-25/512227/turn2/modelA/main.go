package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// LogMessage struct holds log message data
type LogMessage struct {
	Level   string
	Message string
}

// GenerateLogs creates a goroutine to generate log messages
func GenerateLogs(msgChan chan LogMessage, wg *sync.WaitGroup) {
	defer wg.Done()

	logLevels := []string{"INFO", "WARNING", "ERROR", "DEBUG"}
	for {
		// Generate random log level and message
		level := logLevels[rand.Intn(len(logLevels))]
		msg := fmt.Sprintf("Log %d: %s - %s", rand.Intn(1000), level, fmt.Sprintf("Random message %d", rand.Intn(1000)))
		// Send the log message to the channel
		msgChan <- LogMessage{Level: level, Message: msg}
		// Simulate a random interval between log messages
		time.Sleep(time.Duration(rand.Intn(1000)+50) * time.Millisecond)
	}
}

// ProcessLogs is the worker goroutine that processes incoming log messages
func ProcessLogs(msgChan chan LogMessage) {
	for msg := range msgChan {
		fmt.Printf("Processed log: %s: %s\n", msg.Level, msg.Message)
	}
}

func main() {
	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a channel to pass log messages
	msgChan := make(chan LogMessage)

	// WaitGroup to ensure all generator goroutines are completed
	var wg sync.WaitGroup
	const numGenerators = 5 // Number of generator goroutines

	// Start generator goroutines
	for i := 0; i < numGenerators; i++ {
		wg.Add(1)
		go GenerateLogs(msgChan, &wg)
	}

	// Start the worker goroutine
	go ProcessLogs(msgChan)

	// Wait for all generators to complete
	wg.Wait()

	// Gracefully terminate the worker by closing the channel
	close(msgChan)
}
