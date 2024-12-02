package modelA

import (
	"os"

	"github.com/rs/zerolog"
)

var (
	logger *zerolog.Logger
)

// InitLogger initializes the zerolog logger
func InitLogger() {
	logger = zerolog.New(os.Stdout).With().Timestamp().Caller().MsgPrefix("app: ").Log()
}

// Info logs an informational message
func Info(msg string, fields map[string]interface{}) {
	logger.Info().Fields(fields).Msg(msg)
}

// Debug logs a debug message
func Debug(msg string, fields map[string]interface{}) {
	logger.Debug().Fields(fields).Msg(msg)
}

// Error logs an error message
func Error(err error, msg string, fields map[string]interface{}) {
	logger.Error().Err(err).Fields(fields).Msg(msg)
}
