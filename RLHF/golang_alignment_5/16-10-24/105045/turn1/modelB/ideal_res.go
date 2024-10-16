package main

import (
	"fmt"
	"strings"
)

// Define a map of common words to their shorter representations
var compressionMap = map[string]string{
	"the":        "1",
	"is":         "2",
	"for":        "3",
	"and":        "4",
	"to":         "5",
	"a":          "6",
	"sample":     "7",
	"text":       "8",
	"using":      "9",
	"word":       "10",
	"embeddings": "11",
}

func compressText(text string) string {
	// Split the text into words
	words := strings.Fields(text)
	var compressedText []string

	// Replace words using the compression map
	for _, word := range words {
		if replacement, found := compressionMap[strings.ToLower(word)]; found {
			compressedText = append(compressedText, replacement)
		} else {
			compressedText = append(compressedText, word)
		}
	}

	// Join the compressed words back into a single string
	return strings.Join(compressedText, " ")
}

func main() {
	// Example text for compression
	text := "This is a sample text for compression using word embeddings."

	// Compress the text
	compressedText := compressText(text)
	fmt.Println("Original Text:", text)
	fmt.Println("Compressed Text:", compressedText)
}
