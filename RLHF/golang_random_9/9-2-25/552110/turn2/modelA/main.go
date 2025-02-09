package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// initiateSignalHandling sets up a channel to listen for incoming signals
// and handles them in a cross-platform way.
func initiateSignalHandling() {
	// Create a channel to receive incoming signals
	sigs := make(chan os.Signal, 1)

	// Notify the channel on receiving SIGINT (Ctrl+C) and SIGTERM
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine to handle the signals
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				// Handle SIGINT and SIGTERM signals
				fmt.Printf("Received signal: %s, terminating gracefully...\n", sig)
				cleanupAndExit()
			default:
				fmt.Printf("Unhandled signal: %s\n", sig)
			}
		}
	}()
}

// cleanupAndExit performs cleanup tasks before shutting down the application.
func cleanupAndExit() {
	fmt.Println("Performing cleanup operations...")
	// Simulate cleanup, such as closing files or database connections
	// ...
	fmt.Println("Cleanup completed. Exiting now.")
	os.Exit(0)
}

func main() {
	fmt.Println("Application starting... Press Ctrl+C to exit.")

	// Initialize signal handling
	initiateSignalHandling()

	// Simulating application work
	select {} // This will block forever, simulating a long-running application
}
