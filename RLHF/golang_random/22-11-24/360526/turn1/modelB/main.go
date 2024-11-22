package main

import (
	"fmt"
	"log"
)

// ContextualError represents an error with contextual information.
type ContextualError struct {
	err     error
	context string
}

// Error returns the error message with contextual information.
func (e ContextualError) Error() string {
	return fmt.Sprintf("%s: %s", e.context, e.err.Error())
}

// WrapError wraps an existing error with additional contextual information.
func WrapError(err error, context string) error {
	return ContextualError{err: err, context: context}
}

func main() {
	err := fmt.Errorf("an error occurred")
	logError(WrapError(err, "reading file"))
}

func logError(err error) {
	switch v := err.(type) {
	case ContextualError:
		log.Printf("Error: %s\n", v)
	default:
		log.Printf("Error: %s\n", err)
	}
}
