package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	// Example log entry with context variables
	log.WithFields(logrus.Fields{
		"request_id": "123e4567-e89b-12d3-a456-426614174000",
		"user_id":    100,
	}).Info(fmt.Sprintf("Processing request for product ID: %d", 456))
}
