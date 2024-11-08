package main

import (
	"fmt"
	"strings"
)

// isVowel checks if a character is a vowel
func isVowel(c rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, c)
}

// countVowelsInString counts the vowels in a given string
func countVowelsInString(input string) int {
	return len(Filter(input, isVowel))
}

// main function
func main() {
	testStrings := []string{"Hello", "world", "Go is great", "Functional programming rocks!"}

	// Using map to apply the countVowelsInString function to each test string
	vowelCounts := mapString(testStrings, countVowelsInString)

	for _, str := range testStrings {
		fmt.Printf("Vowels in '%s': %d\n", str, vowelCounts[str])
	}
}

// map function implementation
func mapString(strings []string, f func(string) int) map[string]int {
	result := make(map[string]int)
	for _, str := range strings {
		result[str] = f(str)
	}
	return result
}

// strings.Filter implementation for reference
func Filter(s string, f func(rune) bool) string {
	var b strings.Builder
	for _, r := range s {
		if f(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}
