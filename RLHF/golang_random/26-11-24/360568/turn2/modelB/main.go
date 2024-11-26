package main

import (
	"flag"
)

func main() {
	setupLogger()
	logger.Println("Application started")
	logger.Println("Command line arguments:", flag.Args())
	// Application logic goes here
	if debug {
		logger.Println("Debug mode enabled")
	}
	logger.Println("Application completed successfully")
}
