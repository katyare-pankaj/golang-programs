package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// Logger to manage logs
type Logger struct {
	mu     sync.Mutex
	logger *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	return &Logger{
		logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

// Log logs messages with file details and timestamps
func (l *Logger) Log(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.Println(v...)
}

// Decorator function for enhanced error tracking
func withLogging(fn func() error, logger *Logger) func() {
	return func() {
		start := time.Now()
		logger.Log("Starting function execution")
		err := fn()
		if err != nil {
			logger.Log("Error:", err)
		}
		logger.Log("Completed function execution in", time.Since(start))
	}
}

// DummyWorker simulates a worker function
func DummyWorker(id int, logger *Logger) error {
	logger.Log(fmt.Sprintf("Worker %d: Starting work", id))
	time.Sleep(time.Millisecond * time.Duration(100+id*10))
	if id%2 == 0 {
		logger.Log(fmt.Sprintf("Worker %d: Encountered an error!", id))
		return fmt.Errorf("worker %d: simulated error", id)
	}
	logger.Log(fmt.Sprintf("Worker %d: Completed work", id))
	return nil
}

func main() {
	logger := NewLogger()

	// Simulating multiple worker tasks
	taskCount := 5
	var wg sync.WaitGroup

	wg.Add(taskCount)
	for i := 0; i < taskCount; i++ {
		i := i // capture the loop variable
		go func() {
			defer wg.Done()
			workerFunc := withLogging(func() error {
				return DummyWorker(i, logger)
			}, logger)
			workerFunc()
		}()
	}

	wg.Wait()
}
