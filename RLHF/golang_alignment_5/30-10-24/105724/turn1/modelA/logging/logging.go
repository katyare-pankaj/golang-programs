package logging

import "github.com/sirupsen/logrus"

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	// Set up logging configuration as needed
	return logger
}
