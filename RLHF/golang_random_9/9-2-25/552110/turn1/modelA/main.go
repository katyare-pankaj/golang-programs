package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func handleSignals() {
	// Create a channel to receive signals
	sigs := make(chan os.Signal, 1)

	// Notify on specific signals
	// On Linux, you might use syscall.SIGINT and syscall.SIGTERM
	// On Windows, syscall.SIGINT is usually sufficient for Ctrl+C
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine to handle signals
	go func() {
		sig := <-sigs
		fmt.Printf("Received signal: %s\n", sig)
		cleanupAndExit()
	}()
}

func cleanupAndExit() {
	fmt.Println("Performing cleanup...")
	// Perform any necessary cleanup here
	os.Exit(0)
}

func main() {
	fmt.Println("Music player started...")
	handleSignals()

	// Simulate player running
	select {}
}
