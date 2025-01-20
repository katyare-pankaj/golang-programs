package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// SafeLogger struct handles concurrent log writes
type SafeLogger struct {
	mu     sync.Mutex
	logger *log.Logger
}

// NewSafeLogger initializes a SafeLogger
func NewSafeLogger() *SafeLogger {
	return &SafeLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

// Log safely writes log messages
func (sl *SafeLogger) Log(v ...interface{}) {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	sl.logger.Println(v...)
}

// LogWithContext adds context to log messages
func (sl *SafeLogger) LogWithContext(context string, v ...interface{}) {
	v = append([]interface{}{context + ":"}, v...)
	sl.Log(v...)
}

// Decorator to add logging with context
func logDecorator(fn func() error, sl *SafeLogger, context string) func() {
	return func() {
		start := time.Now()
		sl.LogWithContext(context, "Starting execution")
		err := fn()
		if err != nil {
			sl.LogWithContext(context, "Error:", err)
		}
		sl.LogWithContext(context, "Completed execution in", time.Since(start))
	}
}

// ExampleWorker simulates a workload
func ExampleWorker(id int, sl *SafeLogger) error {
	sl.LogWithContext(fmt.Sprintf("Worker-%d", id), "Starting task")
	time.Sleep(time.Millisecond * time.Duration(100+id*10))
	if id%2 == 0 {
		return fmt.Errorf("worker-%d: simulated error", id)
	}
	sl.LogWithContext(fmt.Sprintf("Worker-%d", id), "Task completed")
	return nil
}

func main() {
	sl := NewSafeLogger()

	taskCount := 5
	var wg sync.WaitGroup

	wg.Add(taskCount)
	for i := 0; i < taskCount; i++ {
		i := i // capture the loop variable
		go func(workerID int) {
			defer wg.Done()
			workerFunc := logDecorator(func() error {
				return ExampleWorker(workerID, sl)
			}, sl, fmt.Sprintf("Worker-%d", workerID))
			workerFunc()
		}(i)
	}

	wg.Wait()
}
