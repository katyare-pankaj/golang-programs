package main

import (
	"fmt"
	"os"
)

type Logger interface {
	Info(message string)
	Warn(message string)
	Error(message string)
}
type FileLogger struct {
	file *os.File
}

func NewFileLogger(filename string) (*FileLogger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FileLogger{file: file}, nil
}

func (fl *FileLogger) Info(message string) {
	fmt.Fprintln(fl.file, "INFO:", message)
}

func (fl *FileLogger) Warn(message string) {
	fmt.Fprintln(fl.file, "WARN:", message)
}

func (fl *FileLogger) Error(message string) {
	fmt.Fprintln(fl.file, "ERROR:", message)
}

func (fl *FileLogger) Close() {
	if err := fl.file.Close(); err != nil {
		fmt.Println("Error closing file logger:", err)
	}
}

type ConsoleLogger struct{}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (cl *ConsoleLogger) Info(message string) {
	fmt.Println("INFO:", message)
}

func (cl *ConsoleLogger) Warn(message string) {
	fmt.Println("WARN:", message)
}

func (cl *ConsoleLogger) Error(message string) {
	fmt.Println("ERROR:", message)
}
func main() {
	// Example: Using a FileLogger
	fileLogger, err := NewFileLogger("application.log")
	if err != nil {
		fmt.Println("Error creating file logger:", err)
		return
	}
	defer fileLogger.Close()

	// Example: Using a ConsoleLogger
	consoleLogger := NewConsoleLogger()

	// Use the loggers interchangeably
	loggers := []Logger{fileLogger, consoleLogger}

	for _, logger := range loggers {
		logger.Info("This is an informational message.")
		logger.Warn("This is a warning message.")
		logger.Error("This is an error message.")
	}
}
