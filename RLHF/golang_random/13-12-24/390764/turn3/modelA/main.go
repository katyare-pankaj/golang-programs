package main

import (
	"fmt"
	"log"
	"os"
)

// Logger defines the interface for logging strategies.
type Logger interface {
	Log(message string)
}

// ConsoleLogger is a concrete strategy that logs to the console.
type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("Console: " + message)
}

// FileLogger is a concrete strategy that logs to a file.
type FileLogger struct {
	file *os.File
}

func (f *FileLogger) Log(message string) {
	if f.file != nil {
		log.SetOutput(f.file)
		log.Println("File: " + message)
	}
}

// Close closes the log file.
func (f *FileLogger) Close() {
	if f.file != nil {
		f.file.Close()
	}
}

// LoggerContext is a context to use a logging strategy.
type LoggerContext struct {
	logger Logger
}

// SetLogger sets the logging strategy for the context.
func (c *LoggerContext) SetLogger(logger Logger) {
	c.logger = logger
}

// Log executes the currently set logging strategy.
func (c *LoggerContext) Log(message string) {
	c.logger.Log(message)
}

func main() {
	// Create a LoggerContext
	context := &LoggerContext{}

	// Set the strategy to ConsoleLogger and log a message
	consoleLogger := &ConsoleLogger{}
	context.SetLogger(consoleLogger)
	context.Log("This is a console log.")

	// Set the strategy to FileLogger and log a message
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	fileLogger := &FileLogger{file: file}
	context.SetLogger(fileLogger)
	context.Log("This is a file log.")

	// Close the file logger
	fileLogger.Close()
}
