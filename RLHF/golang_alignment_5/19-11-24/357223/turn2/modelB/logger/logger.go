// logger/logger.go

package logger

import (
	"fmt"
	"log"
)

// LogLevel represents the severity of a log message.
type LogLevel int

const (
	// INFO log level.
	INFO LogLevel = iota
	// WARN log level.
	WARN
	// ERROR log level.
	ERROR
)

var level LogLevel = INFO

// SetLevel sets the log level for the package.
func SetLevel(l LogLevel) {
	level = l
}

// Log logs a message with the specified log level.
func Log(l LogLevel, msg string, args ...interface{}) {
	if l >= level {
		log.Printf(fmt.Sprintf("[%s] %s", levelName(l), msg), args...)
	}
}

func levelName(l LogLevel) string {
	switch l {
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}
