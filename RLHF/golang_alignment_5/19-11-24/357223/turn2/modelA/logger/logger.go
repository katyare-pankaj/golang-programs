// logger.go
package logging

// Logger interface defines logging operations.
type Logger interface {
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
}

// defaultLogger implements Logger interface.
type defaultLogger struct {
}

// Info logs at the INFO level.
func (l *defaultLogger) Info(v ...interface{}) {
	print("INFO: ", v...)
	print("\n")
}

// Warn logs at the WARN level.
func (l *defaultLogger) Warn(v ...interface{}) {
	print("WARN: ", v...)
	print("\n")
}

// Error logs at the ERROR level.
func (l *defaultLogger) Error(v ...interface{}) {
	print("ERROR: ", v...)
	print("\n")
}

// NewLogger returns a new Logger.
func NewLogger() Logger {
	return &defaultLogger{}
}
