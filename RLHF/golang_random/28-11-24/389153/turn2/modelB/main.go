package main

import (
	"fmt"
	"os"
)

type Logger interface {
	Info(message string)
	Error(message string)
	Debug(message string)
}

type FileLogger struct {
	file *os.File
}

func NewFileLogger(filename string) (*FileLogger, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FileLogger{file: file}, nil
}

func (fl *FileLogger) Info(message string) {
	fl.file.WriteString(fmt.Sprintf("INFO: %s\n", message))
}

func (fl *FileLogger) Error(message string) {
	fl.file.WriteString(fmt.Sprintf("ERROR: %s\n", message))
}

func (fl *FileLogger) Debug(message string) {
	fl.file.WriteString(fmt.Sprintf("DEBUG: %s\n", message))
}

type ConsoleLogger struct{}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (cl *ConsoleLogger) Info(message string) {
	fmt.Println("INFO:", message)
}

func (cl *ConsoleLogger) Error(message string) {
	fmt.Println("ERROR:", message)
}

func (cl *ConsoleLogger) Debug(message string) {
	fmt.Println("DEBUG:", message)
}

type LoggerManager struct {
	currentLogger Logger
}

func NewLoggerManager(logger Logger) *LoggerManager {
	return &LoggerManager{currentLogger: logger}
}

func (lm *LoggerManager) SetLogger(logger Logger) {
	lm.currentLogger = logger
}

func (lm *LoggerManager) Info(message string) {
	lm.currentLogger.Info(message)
}

func (lm *LoggerManager) Error(message string) {
	lm.currentLogger.Error(message)
}

func (lm *LoggerManager) Debug(message string) {
	lm.currentLogger.Debug(message)
}
func main() {
	fileLogger, err := NewFileLogger("app.log")
	if err != nil {
		panic(err)
	}

	loggerManager := NewLoggerManager(fileLogger) // Default to file logging

	loggerManager.Info("Application started")
	loggerManager.Debug("Initializing configuration...")

	// Later, you can swap to console logging at runtime
	loggerManager.SetLogger(NewConsoleLogger())
	loggerManager.Error("Something went wrong!")
}
