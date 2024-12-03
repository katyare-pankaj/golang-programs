// logger_test.go
package logger

import (
	"bytes"
	"go-programs/RLHF/golang_random/3-12-24/390900/turn2/modelB/logger"
	"testing"
)

// TestLog verifies that logging works correctly.
func TestLog(t *testing.T) {
	var buf bytes.Buffer
	logger := logger.NewSimpleLogger(&buf)

	expected := "Test message"
	logger.Log(expected)

	if buf.String() != expected+"\n" {
		t.Errorf("Expected %q, got %q", expected, buf.String())
	}
}
