package main

import (
	"fmt"
	"strings"
	"testing"
)

const numStrings = 10000
const strLength = 100

var sampleStrings []string

func init() {
	for i := 0; i < numStrings; i++ {
		sampleStrings = append(sampleStrings, fmt.Sprintf("String %d", i))
	}
}

func concatenateUsingJoin(parts []string) string {
	return strings.Join(parts, " ")
}

func concatenateUsingPlus(parts []string) string {
	result := ""
	for _, part := range parts {
		result += part + " "
	}
	return result[:len(result)-1]
}

func BenchmarkConcatenateUsingJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatenateUsingJoin(sampleStrings)
	}
}

func BenchmarkConcatenateUsingPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatenateUsingPlus(sampleStrings)
	}
}
