package main

import (
	"fmt"
	"strings"
	"testing"
)

// Tokenize is a higher-order function that splits a text into tokens (words).
func Tokenize(text string) []string {
	return strings.Fields(text)
}

// Ngrams is a higher-order function that generates N-grams from a slice of tokens.
func Ngrams(tokens []string, n int) [][]string {
	if n <= 0 {
		return [][]string{}
	}
	ngrams := make([][]string, 0)
	for i := 0; i < len(tokens)-n+1; i++ {
		ngrams = append(ngrams, tokens[i:i+n])
	}
	return ngrams
}

// mapTokens applies a function to each token in a slice and returns a new slice of results.
func mapTokens(tokens []string, f func(string) string) []string {
	result := make([]string, len(tokens))
	for i, token := range tokens {
		result[i] = f(token)
	}
	return result
}

// PartialNgrams returns a function that generates N-grams with a fixed N value.
func PartialNgrams(n int) func([]string) [][]string {
	return func(tokens []string) [][]string {
		return Ngrams(tokens, n)
	}
}

// BenchmarkNgramGeneration benchmarks the performance of N-gram generation.
func BenchmarkNgramGeneration(b *testing.B) {
	text := "Hello, world. This is a test sentence. Let's see how it performs with longer text."
	tokens := Tokenize(text)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Ngrams(tokens, 2)
	}
}

// BenchmarkNgramGenerationWithPartial benchmarks the performance of N-gram generation using partial functions.
func BenchmarkNgramGenerationWithPartial(b *testing.B) {
	text := "Hello, world. This is a test sentence. Let's see how it performs with longer text."
	tokens := Tokenize(text)
	bigramGenerator := PartialNgrams(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bigramGenerator(tokens)
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

	// Example of using mapTokens
	lowercasedTokens := mapTokens(tokens, strings.ToLower)
	fmt.Println("Lowercased Tokens:", lowercasedTokens)
	fmt.Println("Lowercased Bigrams:", bigrams(lowercasedTokens))
}
