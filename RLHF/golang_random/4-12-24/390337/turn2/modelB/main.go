package main

import (
	"fmt"
	"log"
	"os"
)

// Function that could cause a panic
func faultyFunction() {
	// Example panic
	panic("something went wrong!")
}

// Safeguarded function with logging and recovery
func runSafely() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovered from panic: %v", err)
			// Additional handling can be done here
		}
	}()

	faultyFunction()
	fmt.Println("Post-faultyFunction code that won't run if panicked.")
}

func main() {
	logFile, err := os.OpenFile("panic.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	fmt.Println("Starting application...")

	runSafely()

	fmt.Println("Application finished successfully.")
}
