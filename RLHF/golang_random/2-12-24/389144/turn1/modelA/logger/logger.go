package logger

import (
	"io"
)

// LogMessage writes a message to stdout followed by a newline
func LogMessage(writer io.Writer, message string) {
	writer.Write([]byte(message + "\n"))
}
