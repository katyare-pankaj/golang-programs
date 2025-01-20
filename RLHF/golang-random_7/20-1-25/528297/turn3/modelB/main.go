package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// Logger struct to manage logs
type Logger struct {
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
	l.logger.Println(v...)
}

// LogWithContext adds context to log messages
func (l *Logger) LogWithContext(context string, v ...interface{}) {
	v = append([]interface{}{context + ":"}, v...)
	l.Log(v...)
}

// LogErrorWithContext logs errors with additional context
func (l *Logger) LogErrorWithContext(context string, err error, v ...interface{}) {
	v = append([]interface{}{context + ":", err}, v...)
	l.Log(v...)
}

// Decorator to add logging with context and track execution time
func logDecorator(fn func() error, l *Logger, context string) func() {
	return func() {
		start := time.Now()
		l.LogWithContext(context, "Starting execution")
		err := fn()
		if err != nil {
			l.LogErrorWithContext(context, err, "Completed execution with error in")
		} else {
			l.LogWithContext(context, "Completed execution in")
		}
		l.LogWithContext(context, time.Since(start))
	}
}

// DummyWorker simulates a worker function
func DummyWorker(id int, l *Logger) error {
	l.LogWithContext(fmt.Sprintf("Worker-%d", id), "Starting task")
	time.Sleep(time.Millisecond * time.Duration(100+id*10))
	if id%2 == 0 {
		return fmt.Errorf("worker-%d: simulated error", id)
	}
	l.LogWithContext(fmt.Sprintf("Worker-%d", id), "Task completed")
	return nil
}

func main() {
	l := NewLogger()

	taskCount := 5
	var wg sync.WaitGroup

	wg.Add(taskCount)
	for i := 0; i < taskCount; i++ {
		i := i // capture the loop variable
		go func(workerID int) {
			defer wg.Done()
			workerFunc := logDecorator(func() error {
				return DummyWorker(workerID, l)
			}, l, fmt.Sprintf("Worker-%d", workerID))
			workerFunc()
		}(i)
	}

	wg.Wait()
}
