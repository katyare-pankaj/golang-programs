package log

import (
	"log"
)

// LogMessage logs a message prefixed by INFO
func LogMessage(message string) {
	log.Printf("INFO: %s", message)
}

// ErrorMessage logs a message prefixed by ERROR
func ErrorMessage(message string) {
	log.Printf("ERROR: %s", message)
}

// DebugMessage logs a message prefixed by DEBUG
func DebugMessage(message string) {
	log.Printf("DEBUG: %s", message)
}
