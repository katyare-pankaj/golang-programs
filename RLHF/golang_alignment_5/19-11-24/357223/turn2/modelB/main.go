// myapp/main.go

package main

import (
	"fmt"
	"go-programs/RLHF/golang_alignment_5/19-11-24/357223/turn2/modelB/logger"
)

func main() {
	// Set the log level to INFO (default)
	logger.Log(logger.INFO, "Application starting")

	// Log warning message
	logger.Log(logger.WARN, "Memory usage high: %d MB", 120)

	// Set the log level to WARN to filter out INFO messages
	logger.SetLevel(logger.WARN)
	logger.Log(logger.INFO, "Ignored informational message")

	// Log error message
	logger.Log(logger.ERROR, "Fatal error: %v", fmt.Errorf("something went wrong"))
}
