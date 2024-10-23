package main

import (
	"os"
)

func main() {
	// Set GOGC to 70
	os.Setenv("GOGC", "70")
	// Your I/O bound code here
}

// Example using goroutines and channels
func processFiles(filenames []string) {
	ch := make(chan string)
	for _, filename := range filenames {
		go processFile(filename, ch)
	}
	for range filenames {
		<-ch
	}
}

func processFile(filename string, ch chan string) {
	// Process file here
	ch <- filename
}
