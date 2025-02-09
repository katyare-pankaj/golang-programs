package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Handle Ctrl+C (SIGINT) on Windows and Unix
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	go func() {
		for _ = range sigint {
			fmt.Println("Received SIGINT (Ctrl+C), shutting down gracefully...")
			os.Exit(0)
		}
	}()
	// Handle SIGTERM on Unix
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM)
	go func() {
		for _ = range sigterm {
			fmt.Println("Received SIGTERM, shutting down gracefully...")
			os.Exit(0)
		}
	}()
	// Handle SIGBREAK on Windows
	sigbreak := make(chan os.Signal, 1)
	signal.Notify(sigbreak, syscall.SIGBREAK)
	go func() {
		for _ = range sigbreak {
			fmt.Println("Received SIGBREAK, shutting down gracefully...")
			os.Exit(0)
		}
	}()
	fmt.Println("Press Ctrl+C or kill -SIGINT or kill -SIGTERM to stop the program.")
	// Block until one of the signals is received
	select {}
}
