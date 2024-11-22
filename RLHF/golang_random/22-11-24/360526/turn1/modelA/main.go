package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// createLogger creates a configured Zap logger
func createLogger() *zap.Logger {
	core := zapcore.NewDebugCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.NewConsoleSink(),
		zap.AddCallerSkip(1), // skip this function in stack traces
	)
	return zap.New(core)
}

// logError logs an error with contextual information
func logError(ctx context.Context, logger *zap.Logger, msg string, err error) {
	// Create a logger with additional context
	lg := logger.With(
		zap.String("error", msg),
		zap.Error(err),
		zap.String("file", string(extractFilePath(getCallerStack()))),
		zap.Int("line", extractLineNumber(getCallerStack())),
	)

	// Log the error
	lg.Error("Error occurred")
}

// getCallerStack returns the stack trace up to the specified level
func getCallerStack(skip int) []byte {
	// Since the stack trace includes 2 lines for this function itself, and we use zap.AddCallerSkip(1)
	// in createLogger, set skip to 3 to account for the logger function, main function, and the actual caller.
	return zap.Stack(skip)
}

// extractFilePath extracts the file path from the stack trace
func extractFilePath(stack []byte) string {
	return strings.SplitN(string(stack), ";", 1)[0]
}

// extractLineNumber extracts the line number from the stack trace
func extractLineNumber(stack []byte) int {
	parts := strings.SplitN(string(stack), ":", 2)
	if len(parts) < 2 {
		return 0 // Return 0 if the stack trace doesn't include a line number
	}
	lineNumber, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0 // Return 0 if there's an error converting the line number to an integer
	}
	return lineNumber
}

func main() {
	logger := createLogger()
	defer logger.Sync()

	ctx := context.Background()

	// Simulate an error
	err := fmt.Errorf("example error")
	logError(ctx, logger, "An example error occurred", err)
}
