package main

import (
	"flag"
	"log"
	"os"
)

// Define logging levels
const (
	DebugLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	CriticalLevel
)

var (
	logLevel   = flag.Int("log-level", InfoLevel, "Logging level: debug, info, warning, error, critical")
	debugMode  = flag.Bool("debug", false, "Enable debugging mode")
	logger     = log.New(os.Stdout, "", log.LstdFlags)
	levelNames = []string{"debug", "info", "warning", "error", "critical"}
)

func init() {
	flag.Parse()

	// Validate log level
	if *logLevel < DebugLevel || *logLevel > CriticalLevel {
		log.Fatalf("Invalid log level %d. Must be between %d and %d.", *logLevel, DebugLevel, CriticalLevel)
	}

	// Set logger prefix based on debug mode
	if *debugMode {
		logger.SetFlags(log.LstdFlags | log.Lshortfile)
	}
}

func logMessage(level int, format string, args ...interface{}) {
	if level >= *logLevel {
		logger.Printf("[%s] %s %s", levelNames[level], format, args...)
	}
}

func debug(format string, args ...interface{}) {
	logMessage(DebugLevel, format, args...)
}

func info(format string, args ...interface{}) {
	logMessage(InfoLevel, format, args...)
}

func warning(format string, args ...interface{}) {
	logMessage(WarningLevel, format, args...)
}

func error(format string, args ...interface{}) {
	logMessage(ErrorLevel, format, args...)
}

func critical(format string, args ...interface{}) {
	logMessage(CriticalLevel, format, args...)
}

func main() {
	// Example usage of log functions
	debug("Debug information")
	info("Application started")
	warning("A warning occurred")
	error("An error occurred")
	critical("Critical error: shutting down")
}
