// logger_test.go
package logging

import (
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewLogger()

	// Test INFO log level
	logger.Info("This is an info message.")
	// Expected output: INFO: This is an info message.

	// Test WARN log level
	logger.Warn("This is a warn message.")
	// Expected output: WARN: This is a warn message.

	// Test ERROR log level
	logger.Error("This is an error message.")
	// Expected output: ERROR: This is an error message.
}
