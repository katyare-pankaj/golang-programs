package main

import (
	"errors"
	"fmt"
	"log"
	"runtime"
)

// ContextualError holds an error and the context where it occurred.
type ContextualError struct {
	base  error
	ctx   string
	frame runtime.Frame
}

// NewContextualError creates a new ContextualError with the specified base error and context.
func NewContextualError(base error, ctx string) error {
	// Get the runtime frame to capture the call location
	frame := runtime.Caller(1)
	return &ContextualError{base: base, ctx: ctx, frame: frame}
}

// Error implements the error interface to return a formatted error string.
func (ce *ContextualError) Error() string {
	return fmt.Sprintf(
		"%s: %s (file: %s:%d)",
		ce.base.Error(),
		ce.ctx,
		ce.frame.File,
		ce.frame.Line,
	)
}

func main() {
	// Try an operation that might fail
	filePath := "non-existent-file.txt"
	if err := openFile(filePath); err != nil {
		logError(err)
	}
}

func openFile(path string) error {
	// Imagine that this function tries to open a file
	// Normally, you'd have actual file I/O code here
	return NewContextualError(errors.New("file not found"), fmt.Sprintf("trying to open file %s", path))
}

func logError(err error) {
	log.Println(err)
}
