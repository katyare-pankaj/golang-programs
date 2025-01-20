package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

// Logger is a struct to manage logs with additional context
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

// Log logs messages with the provided context
func (l *Logger) Log(ctx context.Context, v ...interface{}) {
	// Extract the context values if available
	caller := runtime.FuncForPC(getCaller(ctx)).Name()
	msg := fmt.Sprint(v...)

	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.Printf("%s: %s\n", caller, msg)
}

// WithContext creates a new context with the provided key-value pairs
func (l *Logger) WithContext(key string, value interface{}) context.Context {
	ctx := context.Background()
	return context.WithValue(ctx, key, value)
}

// getCaller extracts the caller function name from the context
func getCaller(ctx context.Context) uintptr {
	val := ctx.Value("caller")
	if caller, ok := val.(uintptr); ok {
		return caller
	}
	return 0
}

// Decorator function for enhanced error tracking
func withLogging(fn func() error, logger *Logger) func() {
	return func() {
		ctx := logger.WithContext("caller", runtime.Caller(1))
		start := time.Now()
		logger.Log(ctx, "Starting function execution")
		err := fn()
		if err != nil {
			logger.Log(ctx, "Error:", err)
		}
		logger.Log(ctx, "Completed function execution in", time.Since(start))
	}
}

// DummyWorker simulates a worker function
func DummyWorker(id int, logger *Logger) error {
	ctx := logger.WithContext("worker_id", id)
	logger.Log(ctx, fmt.Sprintf("Worker %d: Starting work", id))
	time.Sleep(time.Millisecond * time.Duration(100+id*10))
	if id%2 == 0 {
		logger.Log(ctx, fmt.Sprintf("Worker %d: Encountered an error!", id))
		return fmt.Errorf("worker %d: simulated error", id)
	}
	logger.Log(ctx, fmt.Sprintf("Worker %d: Completed work", id))
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
