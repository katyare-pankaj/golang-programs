package main

import (
	"log"
	"time"
)

type logEntry struct {
	Time     time.Time              `json:"time"`
	Severity string                 `json:"severity"`
	Message  string                 `json:"message"`
	Context  map[string]interface{} `json:"context"`
}

func logStructured(severity, message string, context map[string]interface{}) {
	entry := logEntry{
		Time:     time.Now(),
		Severity: severity,
		Message:  message,
		Context:  context,
	}
	log.Println(entry)
}

func main() {
	// Usage example
	logStructured("INFO", "Application started", map[string]interface{}{
		"version":  "1.0.0",
		"env":      "development",
		"hostname": "example.com",
	})

	logStructured("ERROR", "Database connection failed", map[string]interface{}{
		"error": "connection refused",
	})
}
