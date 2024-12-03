// logger/logger.go
package logger

import (
	"fmt"
	"io"
)

// Logger defines logging behavior with Log method.
type Logger interface {
	Log(message string)
}

// SimpleLogger is a Logger implementation that writes to an io.Writer.
type SimpleLogger struct {
	output io.Writer
}

// NewSimpleLogger creates a new SimpleLogger instance.
func NewSimpleLogger(output io.Writer) *SimpleLogger {
	return &SimpleLogger{output: output}
}

// Log logs a message to the specified output.
func (l *SimpleLogger) Log(message string) {
	fmt.Fprintln(l.output, message)
}
