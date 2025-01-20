package main

import (
	"fmt"
	"log"
	"time"
)

// Logger is an interface for logging messages.
type Logger interface {
	Info(message string)
	Error(message string)
}

// Decorator is a function that wraps a Logger and adds additional behavior.
type Decorator func(Logger) Logger

// NewLogger creates a new Logger with the given decorators.
func NewLogger(decorators ...Decorator) Logger {
	logger := &baseLogger{}
	for _, decorator := range decorators {
		logger = decorator(logger).(*baseLogger)
	}
	return logger
}

// baseLogger is a simple implementation of Logger.
type baseLogger struct{}

func (l *baseLogger) Info(message string) {
	log.Println("INFO:", message)
}

func (l *baseLogger) Error(message string) {
	log.Println("ERROR:", message)
}

// WithTimestamp decorator adds a timestamp to log messages.
func WithTimestamp(logger Logger) Logger {
	return &timestampLogger{logger}
}

type timestampLogger struct {
	logger Logger
}

func (l *timestampLogger) Info(message string) {
	l.logger.Info(fmt.Sprintf("[%s] %s", time.Now().Format(time.RFC3339), message))
}

func (l *timestampLogger) Error(message string) {
	l.logger.Error(fmt.Sprintf("[%s] %s", time.Now().Format(time.RFC3339), message))
}

// WithPrefix decorator adds a prefix to log messages.
func WithPrefix(prefix string) Decorator {
	return func(logger Logger) Logger {
		return &prefixLogger{prefix, logger}
	}
}

type prefixLogger struct {
	prefix string
	logger Logger
}

func (l *prefixLogger) Info(message string) {
	l.logger.Info(l.prefix + message)
}

func (l *prefixLogger) Error(message string) {
	l.logger.Error(l.prefix + message)
}

// main function demonstrates the usage of the logging mechanism.
func main() {
	logger := NewLogger(
		WithTimestamp,
		WithPrefix("APP: "),
	)
	// Use the logger as usual.
	logger.Info("Application started")
	// Simulate some errors.
	for _, err := range []error{fmt.Errorf("error 1"), fmt.Errorf("error 2")} {
		logger.Error(err.Error())
	}
	logger.Info("Application stopped")
}
