package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"
)

// Global variables for flags
var (
	configFile string
	verbose    bool
)

// Parse command-line flags
func parseFlags() {
	flag.StringVar(&configFile, "config", "config.json", "Path to the configuration file")
	flag.BoolVar(&verbose, "v", false, "Enable verbose mode")
	flag.Parse()
}

// Initialize logging
func initLogging() {
	logFormat := "%s\n"
	if verbose {
		logFormat = "[%v] %s\n" // Include timestamp in verbose mode
	}

	// Set custom log format
	log.SetFormatter(func(f *log.Fields) string {
		return strings.Join([]string{
			fmt.Sprintf(logFormat, f["timestamp"], f["message"]),
		}, "")
	})

	// Set log level
	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// Add custom fields to log
	log.SetDefaults(log.Fields{"timestamp": log.FieldFunc(func() interface{} {
		return log.Now().Format("2006-01-02 15:04:05")
	})})
}

func main() {
	// Parse command-line flags
	parseFlags()

	// Initialize logging
	initLogging()

	// Application logic
	log.Infof("Starting application with config file: %s", configFile)
	defer log.Info("Application completed.")

	// Example of debug logging
	if verbose {
		log.Debugf("Debug message: verbose mode is enabled")
	}

	// Simulate some work
	log.Info("Simulating work...")
	select {
	case <-time.After(2 * time.Second):
		log.Info("Work completed successfully.")
	}
}
