package main

import (
	"fmt"
	"io/ioutil"
)

// Logger defines the interface for logging strategies.
type Logger interface {
	Log(message string)
}

// ConsoleLogger is a concrete strategy that logs to the console.
type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("Console:", message)
}

// FileLogger is a concrete strategy that logs to a file.
type FileLogger struct {
	filePath string
}

func NewFileLogger(path string) *FileLogger {
	return &FileLogger{filePath: path}
}

func (f *FileLogger) Log(message string) {
	data := []byte(fmt.Sprintln("File:", message))
	err := ioutil.WriteFile(f.filePath, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

// LogContext is a context to use a logging strategy.
type LogContext struct {
	logger Logger
}

// SetLogger sets the logging strategy for the context.
func (c *LogContext) SetLogger(logger Logger) {
	c.logger = logger
}

// Log executes the currently set logging strategy.
func (c *LogContext) Log(message string) {
	c.logger.Log(message)
}

func main() {
	// Create a logging context
	context := &LogContext{}

	// Set the strategy to ConsoleLogger and log a message
	consoleLogger := &ConsoleLogger{}
	context.SetLogger(consoleLogger)
	context.Log("This is a console log.")

	// Set the strategy to FileLogger and log a message
	fileLogger := NewFileLogger("application.log")
	context.SetLogger(fileLogger)
	context.Log("This is a file log.")
}
