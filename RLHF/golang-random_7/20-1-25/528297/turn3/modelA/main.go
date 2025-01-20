package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// ThreadSafeLogger ensures concurrent access to log writing
type ThreadSafeLogger struct {
	mu     sync.Mutex
	logger *log.Logger
}

// NewThreadSafeLogger creates a new ThreadSafeLogger instance
func NewThreadSafeLogger() *ThreadSafeLogger {
	return &ThreadSafeLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

// Log safely logs messages with optional context
func (tsl *ThreadSafeLogger) Log(context string, v ...interface{}) {
	tsl.mu.Lock()
	defer tsl.mu.Unlock()
	if context != "" {
		v = append([]interface{}{context + ":"}, v...)
	}
	tsl.logger.Println(v...)
}

// LogDecorator decorates a function with logging
func LogDecorator(fn func() error, tsl *ThreadSafeLogger, context string) func() {
	return func() {
		start := time.Now()
		tsl.Log(context, "Start execution")
		err := fn()
		now := time.Now()
		if err != nil {
			tsl.Log(context, "Error:", err)
		}
		tsl.Log(context, "Completed execution in", now.Sub(start))
	}
}

// Worker simulates a workload in the application
func Worker(id int, tsl *ThreadSafeLogger) error {
	context := fmt.Sprintf("Worker-%d", id)
	tsl.Log(context, "Started working")
	time.Sleep(time.Millisecond * time.Duration(100+id*10))

	if id%2 == 0 {
		err := fmt.Errorf("worker-%d encountered an error", id)
		tsl.Log(context, err)
		return err
	}

	tsl.Log(context, "Finished working successfully")
	return nil
}

func main() {
	tsl := NewThreadSafeLogger()

	numWorkers := 5
	var wg sync.WaitGroup

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		i := i // capture loop variable for each goroutine
		go func() {
			defer wg.Done()
			decoratedFunc := LogDecorator(func() error {
				return Worker(i, tsl)
			}, tsl, fmt.Sprintf("Func-Execution:Worker-%d", i))
			decoratedFunc()
		}()
	}

	wg.Wait()
}
