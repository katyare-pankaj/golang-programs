package main

import (
	"flag"
	"log"
	"os"
)

var (
	debug  bool
	logger *log.Logger
)

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debugging mode")
}
func setupLogger() {
	flag.Parse()
	flags := log.LstdFlags
	if debug {
		flags |= log.Lshortfile
	}
	logger = log.New(os.Stdout, "", flags)
}
