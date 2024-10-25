package main

import (
	"fmt"
	"strings"
)

func main() {
	logMessage := "WARNING: This is a warning message with some irrelevant info."
	// Simple text preprocessing: Lowercasing, removing punctuation
	preprocessedMessage := strings.Map(func(r rune) rune {
		if r == '.' || r == ',' || r == '!' || r == '?' {
			return -1
		}
		return r
	}, strings.ToLower(logMessage))
	fmt.Println("Preprocessed Message:", preprocessedMessage)
}
