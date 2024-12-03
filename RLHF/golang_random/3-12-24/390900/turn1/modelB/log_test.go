package log

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestLogMessage(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf) // Redirect log output to buffer for testing
	defer log.SetOutput(nil)

	LogMessage("This is a test log")
	output := buf.String()

	if !strings.Contains(output, "INFO: This is a test log") {
		t.Errorf("Expected 'INFO: This is a test log', got '%s'", output)
	}
}

func TestErrorMessage(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)

	ErrorMessage("This is an error")
	output := buf.String()

	if !strings.Contains(output, "ERROR: This is an error") {
		t.Errorf("Expected 'ERROR: This is an error', got '%s'", output)
	}
}

func TestDebugMessage(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)

	DebugMessage("This is a debug message")
	output := buf.String()

	if !strings.Contains(output, "DEBUG: This is a debug message") {
		t.Errorf("Expected 'DEBUG: This is a debug message', got '%s'", output)
	}
}
