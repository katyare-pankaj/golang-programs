package main

import (
	"fmt"
	"testing"
)

func concatenateStrings(count int) string {
	s := ""
	for i := 0; i < count; i++ {
		s += "Hello, world!"
	}
	return s
}

func sprintfStrings(count int) string {
	return fmt.Sprintf("%s", fmt.Sprintf("%s", "Hello, world!")+fmt.Sprintf("%s", "Hello, world!")+fmt.Sprintf("%s", "Hello, world!")+fmt.Sprintf("%s", "Hello, world!")+fmt.Sprintf("%s", "Hello, world!")+fmt.Sprintf("%s", "Hello, world!")+fmt.Sprintf("%s", "Hello, world!")+fmt.Sprintf("%s", "Hello, world!")+fmt.Sprintf("%s", "Hello, world!")+fmt.Sprintf("%s", "Hello, world!"))
}

func BenchmarkConcatenateStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatenateStrings(1000)
	}
}

func BenchmarkSprintfStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sprintfStrings(1000)
	}
}

func main() {
	testing.Main()
}
