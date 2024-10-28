package main

import (
	"fmt"
	"strings"
)

// Tokenize takes a string and returns a slice of tokens (words)
func Tokenize(text string) []string {
	return strings.Fields(text)
}

// Ngrams generates N-grams from a slice of tokens using a sliding window approach
func Ngrams(tokens []string, n int) [][]string {
	result := make([][]string, 0)
	if n <= 0 {
		return result
	}

	for i := 0; i < len(tokens)-n+1; i++ {
		result = append(result, tokens[i:i+n])
	}
	return result
}

// PartialNgrams returns a function that generates N-grams with a fixed N value
func PartialNgrams(n int) func([]string) [][]string {
	return func(tokens []string) [][]string {
		return Ngrams(tokens, n)
	}
}

func main() {
	text := "Hello, world. This is a test sentence."
	tokens := Tokenize(text)

	// Using Partial N-gram functions to create specific N-gram functions
	unigrams := PartialNgrams(1)
	bigrams := PartialNgrams(2)
	trigrams := PartialNgrams(3)

	// Generate N-grams
	fmt.Println("Unigrams:", unigrams(tokens))
	fmt.Println("Bigrams: ", bigrams(tokens))
	fmt.Println("Trigrams:", trigrams(tokens))
}
