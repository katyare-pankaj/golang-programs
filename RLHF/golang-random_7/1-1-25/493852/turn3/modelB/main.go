package main

import (
	"testing"
)

const numStrings = 10000
const stringLength = 100

func generateStrings() []string {
	strings := make([]string, numStrings)
	for i := 0; i < numStrings; i++ {
		strings[i] = strings.Repeat("a", stringLength)
	}
	return strings
}

func concatenateStrings(strings []string) string {
	var result string
	for _, str := range strings {
		result += str
	}
	return result
}

func joinStrings(strings []string) string {
	return strings.Join(strings, "")
}

func BenchmarkConcatenation(b *testing.B) {
	strings := generateStrings()
	for i := 0; i < b.N; i++ {
		concatenateStrings(strings)
	}
}

func BenchmarkJoin(b *testing.B) {
	strings := generateStrings()
	for i := 0; i < b.N; i++ {
		joinStrings(strings)
	}
}
