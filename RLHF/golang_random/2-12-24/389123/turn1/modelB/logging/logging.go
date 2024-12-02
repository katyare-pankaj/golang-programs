package logging

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	// Add other log levels as needed
}
