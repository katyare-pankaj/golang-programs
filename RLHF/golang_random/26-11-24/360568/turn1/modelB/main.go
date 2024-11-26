package main

import (
	"flag"
	"log"
	"os"
)

var (
	debug      bool
	configFile string
	logger     = log.New(os.Stdout, "", log.LstdFlags)
)

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debugging mode")
	flag.StringVar(&configFile, "config", "config.json", "configuration file path")
	flag.Parse()
	if debug {
		logger.SetFlags(log.LstdFlags | log.Lshortfile)
	}
}
func main() {
	logger.Println("Application started")
	logger.Println("Loading configuration from:", configFile)
	// Load configuration from file
	// Perform application tasks using configuration settings
	logger.Println("Application completed successfully")
}
