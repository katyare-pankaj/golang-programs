package main

import (
	"encoding/json"
	"log"
	"strings"
	"testing"
	"time"
)

func TestLogStructured(t *testing.T) {
	// Capture the log output for testing
	var logOutput strings.Builder
	log.SetOutput(&logOutput)
	defer log.SetOutput(nil) // Restore default log output

	// Test the INFO log entry
	logStructured("INFO", "Application started", map[string]interface{}{
		"version":  "1.0.0",
		"env":      "development",
		"hostname": "example.com",
	})

	// Parse the logged JSON entry
	var infoEntry logEntry
	if err := json.Unmarshal([]byte(logOutput.String()), &infoEntry); err != nil {
		t.Fatalf("Error parsing log JSON: %v", err)
	}

	// Verify the INFO log entry fields
	expectedInfoEntry := logEntry{
		Time:     time.Now().Round(time.Second), // Round to ignore nanoseconds
		Severity: "INFO",
		Message:  "Application started",
		Context: map[string]interface{}{
			"version":  "1.0.0",
			"env":      "development",
			"hostname": "example.com",
		},
	}
	if infoEntry != expectedInfoEntry {
		t.Errorf("INFO log entry mismatch: expected %v, got %v", expectedInfoEntry, infoEntry)
	}

	// Test the ERROR log entry
	logOutput.Reset()
	logStructured("ERROR", "Database connection failed", map[string]interface{}{
		"error": "connection refused",
	})

	// Parse the logged JSON entry
	var errorEntry logEntry
	if err := json.Unmarshal([]byte(logOutput.String()), &errorEntry); err != nil {
		t.Fatalf("Error parsing log JSON: %v", err)
	}

	// Verify the ERROR log entry fields
	expectedErrorEntry := logEntry{
		Time:     time.Now().Round(time.Second), // Round to ignore nanoseconds
		Severity: "ERROR",
		Message:  "Database connection failed",
		Context: map[string]interface{}{
			"error": "connection refused",
		},
	}
	if errorEntry != expectedErrorEntry {
		t.Errorf("ERROR log entry mismatch: expected %v, got %v", expectedErrorEntry, errorEntry)
	}
}
